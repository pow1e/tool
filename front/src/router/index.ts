/*
 * @Description: 
 * @version: 
 * @Author: William
 * @Date: 2023-07-12 15:27:12
 * @LastEditors: William
 * @LastEditTime: 2023-07-25 10:46:10
 */
import { createRouter, createWebHashHistory } from "vue-router";

import Layout from "../views/layout/index.vue";
import Index from "../views/index/index.vue";
import QuickCommand from "../views/quickCommand/index.vue";
import md5Translate from "../views/md5Translate/index.vue";
import encrypt from "../views/md5Translate/components/encrypt/index.vue";
import decrypt from "../views/md5Translate/components/decrypt/index.vue";
const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      component: Layout,
      children: [
        {
          path: "/",
          component: Index
        },
        {
          path: "/qc",
          component: QuickCommand
        },
        {
          path: "/md5",
          component: md5Translate,
          children: [
            {
              path: "/encrypt",
              component: encrypt
            },
            {
              path: "/decrypt",
              component: decrypt
            }
          ]
        }
      ]
    }
  ]
});

export default router