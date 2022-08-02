<template>
    <div id="app" class="text-center">
        <Nav />
        <br>
        <br>
        <h1><b>UPLOAD</b></h1>
        <br>
        <HelloWorld id="hello" v-on:uuid="name = $event" />
        <br>
        <br>
        <br>
        <h1><b>RESULT</b></h1>
        <br>
        <button onclick="window.location.reload();" class="btn btn-outline-success">結果を表示</button>
        <br>
        <br>
        <img :src="require(`@/data/${this.imgName}.png`)" height=250>
        <br>
        <br>
        <p>30秒後を目安に"結果を表示"ボタンを押してみてください</p>
    </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import Nav from "./components/Nav.vue"

export default {
    name: 'App',
    data: function () {
        return {
            name: String,
            imgName: "output",
        }
    },
    computed: {
        imgSrc() {
            return require("./data/" + this.imgName + "_output.png")
        },
    },
    methods: {
        changeImages: function () {
            this.imgName = "output"
        }
    },
    components: {
        HelloWorld,
        Nav
    },
}
</script>

<style>
#hello {
    margin-left: 100px;
    margin-right: 100px;
}
</style>
