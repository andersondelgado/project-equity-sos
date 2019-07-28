import Vue from 'vue';
import Router from 'vue-router';
// import Home from './views/Home.vue';
import About from './views/About.vue';
import Login from './components/Login.vue';
import Register from './components/Register.vue';
import Index from './components/Index.vue';
import Test from './components/Test.vue';
import Post from './components/Post.vue';
import BusinessPost from './components/BusinessPost.vue';
import AddPost from './components/AddPost.vue';
import UserPost from './components/UserPost.vue';
import Article from './components/Article.vue';
import Category from './components/Category.vue';
import RegisterUsers from './components/RegisterUsers.vue';

import store from './store';
Vue.use(Router);

const routes = [
  {
    path: '/',
    redirect: {
      name: 'main'
    }
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
  },
  {
    path: '/register',
    name: 'register',
    component: Register,
  },
  {
    path: '/main',
    name: 'main',
    component: Index,
    children: [
      {
        path: '',
        component: About
      },
      {
        path: 'test',
        name: 'test',
        component: Test,
        meta: { requiresPer: true }
      },
      {
        path: 'post',
        name: 'post',
        component: Post,
        meta: { requiresPer: true }
      },
      {
        path: 'add-post',
        name: 'add-post',
        component: AddPost,
        meta: { requiresPer: true }
      },
      {
        path: 'business-post',
        name: 'business-post',
        component: BusinessPost,
        meta: { requiresPer: true }
      },
      {
        path: 'user-post',
        name: 'user-post',
        component: UserPost,
        meta: { requiresPer: true }
      },
      {
        path: 'article',
        name: 'article',
        component: Article,
        meta: { requiresPer: true }
      },
      {
        path: 'category',
        name: 'category',
        component: Category,
        meta: { requiresPer: true }
      },
      {
        path: 'register-users',
        name: 'register-users',
        component: RegisterUsers,
        meta: { requiresPer: true }
      }
    ],
    meta: { requiresAuth: true },
  },
  // {
  //   path: '/index',
  //   name: 'index',
  //   component: Index,
  //   meta: { requiresAuth: true },
  // }
];

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

// const router = new Router({
//   base: process.env.BASE_URL,
//   routes,
// });

router.beforeEach((to, from, next) => {
  // console.log('#to: '+JSON.stringify(to)+' #from: '+JSON.stringify(from)+' #next: '+next);
  // check if the route requires authentication and user is not logged in
  if (to.matched.some((route) => route.meta.requiresAuth) && !store.state.isLoggedIn) {
    // redirect to login page
    next({ path: '/login' });
    return;
  }
  // 
  if (to.matched.some(route => route.meta.requiresPer)) {

    // console.log("# f(to.path): " + f(to.path));
    if (f(to.path) === to.meta.requiresPer) {
      // console.log("hola path: " + to.path + " to.meta.requiresPer: " + to.meta.requiresPer);
      next()
      return
    } else {
      // console.log("to main");
      next({ path: '/main' })
      return
    }
  }

  if (to.path === '/login' && store.state.isLoggedIn) {
    next({ path: '/main' });
    return;
  }

  if (to.path === '/register' && store.state.isLoggedIn) {
    next({ path: '/main' });
    return;
  }

  function f(URL: any) {

    var isTrue = false;
    let per: any = localStorage.getItem('permission');
    if (!per) { isTrue = false; } else {
      let data: any = JSON.parse(per);
      // const data=(localStorage.getItem('permission'));
      for (var i = 0; i < data.length; i++) {
        if (data[i].url === URL) {
          isTrue = true;
        }
      }
    }
    return isTrue;
  };

  next();

});

export default router;

// export default new Router({
//   mode: 'history',
//   base: process.env.BASE_URL,
//   routes: [
//     {
//       path: '/',
//       name: 'home',
//       component: Home,
//     },
//     {
//       path: '/about',
//       name: 'about',
//       // route level code-splitting
//       // this generates a separate chunk (about.[hash].js) for this route
//       // which is lazy-loaded when the route is visited.
//       component: () => import(/* webpackChunkName: "about" */ './views/About.vue'),
//     },
//   ],
// });
