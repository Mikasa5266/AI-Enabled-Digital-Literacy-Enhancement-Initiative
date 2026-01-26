import Home from "../views/Home.vue";
import CourseCreat from "../views/Resource/CourseCreat.vue";
import PPT from "../views/Resource/PPT.vue";
import AIPrepare from "../views/Teacher/AIPrepare.vue";
import AIRevise from "../views/Teacher/AIRevise.vue";
import EverydayWork from "../views/Interact/EverydayWork.vue";
import { createRouter, createWebHistory } from "vue-router";
import { Components } from "ant-design-vue/es/date-picker/generatePicker";
import WorkShow from "../views/Interact/WorkShow.vue";

const routes = [
    {
        path:'/Coursecreat',
        name:'Coursecreat',
        component:CourseCreat
    },
    {
        path:'/PPT',
        name:'PPT',
        component:PPT
    },
    {
        path:'/AIPrepare',
        name:'AIPrepare',
        component:AIPrepare
    },
    {
        path:'/AIRevise',
        name:'AIRevise',
        component:AIRevise
    },
    {
        path:'/EverydayWork',
        name:'EverydayWork',
        component:EverydayWork
    },
    {
        path:'/WorkShow',
        name:'WorkShow',
        component:WorkShow
    }
]
const router = createRouter({
    history:createWebHistory(),
    routes
})
export default router