<template>
  <div id="app" class="text-center">
    <Nav />
    <br>
    <br>
    <h1><b>UPLOAD</b></h1>
    <br>
    <HelloWorld id="hello" />
    <br>
    <br>
    <br>
    <h1><b>RESULT</b></h1>
    <img src="./data/7b28e820-3127-43e7-86aa-5fc5b1fe873c_output.png" alt="">
    <br>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import Nav from "./components/Nav.vue"

export default {
  name: 'App',
  components: {
    HelloWorld,
    Nav
  }
}
</script>

<style>
#hello {
    margin-left:  100px;
    margin-right:  100px;
}
</style>
