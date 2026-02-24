<template>
    <div>
        <Navbar/>
        <div style="padding:20px">
            <router-view/>
        </div>
    </div>
</template>

<script setup>
import Navbar from '../components/Navbar.vue';

</script>