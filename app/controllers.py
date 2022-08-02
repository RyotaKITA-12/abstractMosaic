import skimage.util
import skimage.transform
import skimage.io
import numpy as np
import matplotlib.pyplot as plt
from fastapi import FastAPI
from pydantic import BaseModel
from starlette.requests import Request
import glob
import random

import matplotlib
matplotlib.use("Agg")

app = FastAPI()


class Path(BaseModel):
    path: str


def index(request: Request):
    return {'Hello': 'World'}


def trim_image(height_mosaic, width_mosaic, img_block):
    height_block = img_block.shape[0]
    width_block = img_block.shape[1]
    ratio_mosaic = width_mosaic / height_mosaic
    ratio_block = width_block / height_block
    if ratio_block == ratio_mosaic:
        img_trimmed = skimage.transform.resize(
            img_block,
            (height_mosaic, width_mosaic, 3),
            anti_aliasing=True)
    elif ratio_block > ratio_mosaic:
        img_trimmed = skimage.transform.resize(
            img_block,
            (height_mosaic, np.floor(ratio_block*height_mosaic), 3),
            anti_aliasing=True)
        img_trimmed = img_trimmed[:, :width_mosaic]
    else:
        img_trimmed = skimage.transform.resize(
            img_block,
            (np.floor(width_mosaic*height_block/width_block), width_mosaic, 3),
            anti_aliasing=True)
        img_trimmed = img_trimmed[:height_mosaic, :]

    return img_trimmed


def resized_lab_img(img):
    return skimage.color.rgb2lab(skimage.transform.resize(img,
                                                          (8, 8, 3),
                                                          anti_aliasing=True))


def create_mosaic(p: Path):
    img_mosaic = skimage.io.imread(p.path)

    row = 40
    col = 40

    img = skimage.transform.resize(
        img_mosaic, (img_mosaic.shape[0]*2,
                     img_mosaic.shape[1]*2, 3),
        anti_aliasing=True)
    row_size = img.shape[0] // row
    col_size = img.shape[1] // col
    img_mosaic_trimmed = img[:(row_size*row), :(col_size*col), :]
    img_mosaics = skimage.util.view_as_blocks(
        img_mosaic_trimmed, (row_size, col_size, 3))

    paths = glob.glob('./data/abstract_art/*.jpg')
    paths_rand_300 = random.sample(paths, 301)

    n = 0
    indexes_path = []
    img_blocks = np.zeros((row_size, col_size, 3, 300))
    for i, path in enumerate(paths_rand_300):
        if i == 300:
            break
        img_path = skimage.io.imread(path)
        if len(img_path.shape) == 2:
            continue
        indexes_path.append(i)
        img_blocks[:, :, :, n] = trim_image(row_size, col_size, img_path)
        n += 1

    img_blocks = img_blocks[:, :, :, :n]

    img_lab_blocks = np.zeros((8, 8, 3, n))
    for i in range(n):
        img_lab_blocks[:, :, :, i] = resized_lab_img(img_blocks[:, :, :, i])

    img_M = np.zeros((row_size*row, col_size*col, 3))
    for r in range(row):
        for c in range(col):
            img_m = resized_lab_img(img_mosaics[r, c, 0, :, :, :])
            diff_min = np.sum(skimage.color.deltaE_cie76(
                img_m, img_lab_blocks[:, :, :, 0]))
            index_min = 0
            for i in range(1, n):
                diff = np.sum(skimage.color.deltaE_cie76(
                    img_m, img_lab_blocks[:, :, :, i]))
                if diff < diff_min:
                    diff_min = diff
                    index_min = i
            img_M[(row_size*r):((r+1)*row_size),
                  (col_size*c):((c+1)*col_size),
                  :] = img_blocks[:, :, :, index_min]

    fig = plt.figure()
    plt.imshow(img_M)
    plt.axis("off")
    # path_out = p.path[:-4] + "_output.png"
    path_out = "/tmp/share/output.png"
    fig.savefig(path_out, transparent=True)
    return {"res": "ok", "input_path": p.path, "output_path": path_out}
