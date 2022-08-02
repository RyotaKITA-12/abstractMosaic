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
        }
    },
    computed: {
        imgSrc() {
            return require("./data/" + this.name + "_output.png")
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
