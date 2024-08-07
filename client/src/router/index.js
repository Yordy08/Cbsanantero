// router.js
import { createRouter, createWebHashHistory } from 'vue-router';
import HomeView from '../views/HomeView.vue';
import LoginView from '../views/LogincbView.vue';

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView,
  },
  {
    path: '/about',
    name: 'about',
    component: () => import('../views/AboutView.vue'),
  },
  {
    path: '/principal',
    name: 'princ',
    component: () => import('../views/PrincView.vue'),
  },
  {
    path: '/artes',
    name: 'art',
    component: () => import('../views/ArteView.vue'),
  },
  {
    path: '/bard',
    name: 'bar',
    component: () => import('../views/BardView.vue'),
  },
  {
    path: '/hotels',
    name: 'hote',
    component: () => import('../views/HotelsView.vue'),
  },
  {
    path: '/hospe',
    name: 'hopeda',
    component: () => import('../views/HospeView.vue'),
  },
  {
    path: '/recre',
    name: 'recrat',
    component: () => import('../views/RecreaView.vue'),
  },
  {
    path: '/tourn',
    name: 'tures',
    component: () => import('../views/TourView.vue'),
  },
  {
    path: '/tranport',
    name: 'transp',
    component: () => import('../views/TransporView.vue'),
  },
  {
    path: '/restaurantes',
    name: 'res',
    component: () => import('../views/RestauranteView.vue'),
  },
  {
    path: '/admin',
    name: 'ad',
    component: () => import('../views/AdminView.vue'),
  },
  {
    path: '/login',
    name: 'log',
    component: LoginView,
  },
  {
    path: '/usuario',
    name: 'us',
    component: () => import('../views/listar/ListarUsuario.vue'),
  },
  {
    path: '/lisart',
    name: 'ar',
    component: () => import('../views/listar/ListarArtesanias.vue'),
  },
  {
    path: '/lishospe',
    name: 'ho',
    component: () => import('../views/listar/ListarHospedaje.vue'),
  },
  {
    path: '/listarhoteles',
    name: 'te',
    component: () => import('../views/listar/ListarHotel.vue'),
  },
  {
    path: '/listarTrasporte',
    name: 'ta',
    component: () => import('../views/listar/ListarTrasporte.vue'),
  },
  
  {
    path: '/listarbar',
    name: 'ba',
    component: () => import('../views/listar/ListarBar.vue'),
  },
  {
    path: '/listjuegos',
    name: 'jue',
    component: () => import('../views/listar/ListarJuegos.vue')
  },
  {
    path: '/listtaurant',
    name: 'Rest',
    component: () => import('../views/listar/ListarRestaurant.vue'),
  },
  {
    path: '/artegaleria',
    name: 'Galeri',
    component: () => import('../views/galeria/GaleriVew.vue')
  },
  {
    path: '/resgister',
    name: 'Res',
    component: () => import('../views/RestroVew.vue')
  },
  {
    path: '/notification',
    name: 'noti',
    component: () => import('../views/notification/NotifiView.vue')
  }
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
