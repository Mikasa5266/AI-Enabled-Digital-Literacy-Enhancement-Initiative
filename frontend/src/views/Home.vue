<template>
    <div class="home">
        <aside :class="['home-nav', theme === 'dark' ? 'home-nav-dark' : 'home-nav-light']">
            <div>
                <Switch :checked="theme == 'dark'" checked-children="Dark" un-checked-children="Light"
                    @change="changetheme" />
                <br />
                <br />
                <Menu v-model:open-keys="openkeys" v-model:selected-keys="selectedkeys" mode="inline" :theme="theme"
                    :items="items" @click="handleMenuClick" />
            </div>
        </aside>
        <main class="home-content">
            <RouterView />
        </main>
    </div>
</template>

<script lang="ts" setup>
import { Menu, Switch, type MenuProps, type MenuTheme } from 'ant-design-vue';
import { ref, watchEffect } from 'vue';
import { useRoute, useRouter } from 'vue-router';
//路由控制
const router = useRouter()
const route = useRoute()
//主题切换
const THEME_STORAGE_KEY = 'ai-vue-app-theme'
const theme = ref<MenuTheme>(localStorage.getItem(THEME_STORAGE_KEY) as MenuTheme || 'dark')
const changetheme = (checked: string | number | boolean) => {
    const newTheme = checked ? 'dark' : 'light'
    theme.value = newTheme
    localStorage.setItem(THEME_STORAGE_KEY, newTheme)
}
//页面内容导航
const openkeys = ref<string[]>([])
const selectedkeys = ref<string[]>([])
const items = ref([
    {
        key: '1',
        label: '资源库',
        title: '资源库',
        children: [
            {
                key: '/PPT',
                label: 'PPT',
                title: 'PPT',
            },
            {
                key: '/Coursecreat',
                label: 'AI课程创作',
                title: 'AI课程创作',
            }
        ]
    },
    {
        key: '2',
        label: '教师支持专区',
        title: '教师支持专区',
        children: [
            {
                key: '/AIPrepare',
                label: 'AI 备课工具使用指南',
                title: 'AI 备课工具使用指南',
            },
            {
                key: '/AIRevise',
                label: 'AI批改作业示例',
                title: 'AI批改作业示例',
            }
        ]
    },
    {
        key: '3',
        label: '互动学习区',
        title: '互动学习区',
        children: [
            {
                key: '/EverydayWork',
                label: '作品展示墙',
                title: '作品展示墙',
            },
            {
                key: '/WorkShow',
                label: '每日任务',
                title: '每日任务',
            }
        ]
    }
])

const handleMenuClick: MenuProps['onClick'] = ({ key }) => {
    router.push(key as string)
    console.log(key)
}

//监听窗口刷新
watchEffect(() => {
    selectedkeys.value = [route.path]
    items.value.forEach(parent => {
        if (parent.children) {
            const isChildActive = parent.children.some(chile => chile.key === route.path)
            if (isChildActive) {
                if (!openkeys.value.includes(parent.key)) {
                    openkeys.value = [parent.key]
                }
            }
        }
    });
})
</script>

<style scoped>
.home {
    display: flex;
    height: 100vh;
    width: 100vw;
    overflow: hidden;
    justify-content: left;
}

.home-nav {
    width: 260px;
    flex-shrink: 0;
}

.home-nav-dark {
    background-color: #001529;
    color: white;
}

.home-nav-light {
    background-color: #fcfcfc;
    color: rgba(0, 0, 0, 0.85);
}

.home-content {
    flex-grow: 1;
    overflow-y: auto;
    height: 100vh;
}
</style>