<template>
    <div class="custom-layout" :style="{gap,justifyContent,alignItems,flexDirection: inline ? 'row' : 'column'}" >
        <slot/>
    </div>
</template>
<script setup lang="ts">
    const props = defineProps({
        gap:{
            type:String,
            default:'16px'
        },
        justifyContent:{
            type:String
        },
        alignItems:{
            type:String
        },
        inline:{
            type:String
        }
    })
</script>
<style>
    .custom-layout{
        display: flex;
        h1,h2,h3,h4,h5,h6,p{
            margin:0;
        }
    }
</style>