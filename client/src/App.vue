
<template>
    <!-- eslint-disable -->
  <nav>
    <!-- Topbar Start -->
    <div class="container-fluid bg-light pt-3 d-none d-lg-block">
      <div class="container">
        <div class="row">
          <div class="col-lg-6 text-center text-lg-left mb-2 mb-lg-0">
            <div class="d-inline-flex align-items-center">
              <p><i class="fa fa-envelope mr-2"></i>costabrisaweb@gmail.com</p>
              <p class="text-body px-3">|</p>
              <p><i class="fa fa-phone-alt mr-2"></i>+57 312 5637781</p>
            </div>
          </div>
          <div class="col-lg-6 text-center text-lg-right">
            <div class="d-inline-flex align-items-center" id="res">
              <a class="text-primary px-3" href="https://www.facebook.com/profile.php?id=100087514673618">
                <i class="fab fa-facebook-f"></i>
              </a>
              <a class="text-primary px-3" href="https://www.instagram.com/_costabrisa/">
                <i class="fab fa-instagram"></i>
              </a>
             

             <!--  <a class="text-primary pl-3" href="">
                <i class="fab fa-rr-bell"></i>
              </a> -->
              
              <router-link class="text-primary pl-3" v-if="usuarioAutenticado && usuarioRol === 'Admin' || usuarioAutenticado && usuarioRol === 'Vendedor'" to="/admin">Administrador</router-link>
              
              <router-link class="text-primary pl-3" v-if="!usuarioAutenticado"  to="/login">Iniciar</router-link>
              <div  class="nav-item nav-link">

               
               
                <i  v-if="usuarioAutenticado && usuarioRol !== ''" @click="cerrarSesion()" class="fas fa-sign-out-alt"></i> 
              
              </div>
  
              <div v-if="usuarioAutenticado" class="usuario-info">
                <img :src="imagenUsuario" alt="Bienvenido" loading="lazy" class="avatar" >
                <p class="nombre-usuario">{{ nombreUsuario }}</p>
              </div>
                 <p style="color: transparent;">..</p>
                 <div class="notification-container">
                
              </div>
               
              <router-link to="/notification"> 
              <button  v-if="usuarioRol === 'Vendedor' || usuarioRol === 'Admin'"  type="button" class="btn btn-primary position-relative">
                    <i class="fas fa-bell"></i>
                    <span v-if="notificaciones.length > 0" class="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-danger">
                      {{ notificaciones.length}}
                     <!--  <span class="visually-hidden">Solicitud..</span> --><!--  -->
                    
                    </span>
                  </button>
                </router-link>

            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Topbar End -->

    <!-- Navbar Start -->
    <div class="container-fluid position-relative nav-bar p-0">
      <div class="container-lg position-relative p-0 px-lg-3" style="z-index: 9;">
        <nav class="navbar navbar-expand-lg bg-light navbar-light shadow-lg py-3 py-lg-0 pl-3 pl-lg-5">
          <a href="" class="navbar-brand">
            <h1 class="m-0 text-primary"><span class="text-dark">COSTA</span>BRISA</h1>
          </a>
          <button type="button" class="navbar-toggler" data-toggle="collapse" data-target="#navbarCollapse">
            <span class="navbar-toggler-icon"></span>
          </button>
          <div class="collapse navbar-collapse justify-content-between px-3" id="navbarCollapse">
            <div class="navbar-nav ml-auto py-0">
              <!--
                  <router-link v-if="usuarioAutenticado && usuarioRol === 'admin'" to="/admin">Admin</router-link>
              <router-link v-if="usuarioAutenticado && usuarioRol === 'admin'" to="/admin">Publicar</router-link>
              -->
            
              <div class="nav-item d-lg-none">
                <router-link to="/notification" @click="ocultarMenu"> 
                
              <button   v-if="usuarioAutenticado && usuarioRol === 'Vendedor' || usuarioRol === 'Admin'"  type="button" class="btn btn-primary position-relative">
                    <i class="fas fa-bell"> Notificaciones</i>
                    <span v-if="notificaciones.length > 0" class="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-danger">
                      {{ notificaciones.length}}
                     <!--  <span class="visually-hidden">Solicitud..</span> --><!--  -->
                    
                    </span>
                  </button>
                </router-link>
            </div>

              
            <div class="nav-item d-lg-none">
              <router-link @click="ocultarMenu" v-if="!usuarioAutenticado" to="/login" class="nav-link">
                <i  class="fas fa-sign-in-alt"></i> Iniciar sesión
              </router-link>
            </div>
           
           


              <a href="" class="nav-item nav-link">
                <router-link @click="ocultarMenu" to="/">
                   Inicio
                </router-link>
              </a>
              <router-link @click="ocultarMenu" to="/about" class="nav-item nav-link">
               Acerca de
              </router-link>
           <!--    <a href="service.html" class="nav-item nav-link">
               Mapa turístico
               </a> -->
               
               <router-link @click="ocultarMenu" to="/tranport" class="nav-item nav-link">
                Transporte
                </router-link>

              
              <div class="nav-item dropdown">
                <a  href="#" class="nav-link dropdown-toggle" data-toggle="dropdown">
                 Servicios
                </a>

                <div class="dropdown-menu border-0 rounded-0 m-0">
                 <!--  <router-link @click="ocultarMenu" to="/artes" class="dropdown-item">Atesanias</router-link> -->
                 <router-link @click="ocultarMenu" to="/hotels" class="dropdown-item">Zona Hotelera</router-link>
<!--                   <router-link @click="ocultarMenu" to="/tranport" class="dropdown-item">Transporte</router-link>
 -->              <!--  <router-link @click="ocultarMenu" to="/hospe" class="dropdown-item">Hospedaje</router-link> -->
                 
                  <router-link @click="ocultarMenu" to="/restaurantes" class="dropdown-item">Restaurantes</router-link>
                  <!-- <router-link @click="ocultarMenu" to="/bard" class="dropdown-item">Bares</router-link> -->
                  <!-- <router-link @click="ocultarMenu" to="/recre" class="dropdown-item">Juegos y Recreación</router-link> -->
                </div>
              </div>

               

              <!-- <a href="contact.html" class="nav-item nav-link">
                <i class="fas fa-envelope"></i> Contact
              </a> -->

             <a class="nav-item nav-link" href="">
                  <div class="nav-item d-lg-none">
                  <router-link @click="ocultarMenu" class="text-primary pl-3" v-if="usuarioAutenticado && usuarioRol === 'Admin' || usuarioAutenticado && usuarioRol === 'Vendedor'" to="/admin">Administrador</router-link>
                </div>
             </a>

              <div class="nav-item d-lg-none">
              <i   v-if="usuarioAutenticado && usuarioRol !== ''" @click="cerrarSesion()" class="fas fa-sign-out-alt">Cerrar</i> 
            </div>
          

            </div>
          </div>
        </nav>
      </div>
    </div>
    <!-- Navbar End -->

    <router-view></router-view>
  </nav>
  


  

    <!-- Footer Start -->
    <div class="container-fluid bg-dark text-white-50 py-5 px-sm-3 px-lg-5" style="margin-top: auto;">
        <div class="row pt-6">
            <div class="col-lg-3 col-md-6 mb-5">
                <a href="" class="navbar-brand">
                    <h1 class="text-primary"><span class="text-white">COSTA</span>BRISA</h1>
                </a>
                <p>Es un proyecto  orientado a fortalecer el reconocimiento de diversos sectores atractivos de la Costa.</p>
                <h6 class="text-white text-uppercase mt-4 mb-3" style="letter-spacing: 5px;">Siguenos en</h6>
                <div class="text-white text-uppercase mt-4 mb-3">
                 <!--    <a class="btn btn-outline-primary btn-square mr-2" href="#"><i class="fab fa-twitter"></i></a> -->
                    <a class="btn btn-outline-primary btn-square mr-2" href="https://www.facebook.com/profile.php?id=100087514673618"><i class="fab fa-facebook-f"></i></a>
                    <!-- <a class="btn btn-outline-primary btn-square mr-2" href="#"><i class="fab fa-linkedin-in"></i></a> -->
                    <a class="btn btn-outline-primary btn-square" href="https://www.instagram.com/_costabrisa/"><i class="fab fa-instagram"></i></a>
                </div>
            </div>
            <div class="col-lg-3 col-md-6 mb-5">
                <h5 class="text-white text-uppercase mb-4" style="letter-spacing: 5px;">Servicios</h5>
                <div class="d-flex flex-column justify-content-start">
                  <router-link class="text-white-50 mb-2" to="/tranport" @click.native="handleClick">
                    <i class="fa fa-angle-right mr-2"></i>Transporte
                  </router-link>
                   
                  <router-link class="text-white-50 mb-2" to="/hotels" @click.native="handleClick">
                    <i class="fa fa-angle-right mr-2"></i>Zona Hotelera
                  </router-link>

                  <router-link class="text-white-50 mb-2" to="/restaurantes" @click.native="handleClick">
                    <i class="fa fa-angle-right mr-2"></i>Restaurantes
                  </router-link>

                
                    
                
            
                   <!--  <a class="text-white-50 mb-2" href="#"><i class="fa fa-angle-right mr-2"></i>Packages</a>
                    <a class="text-white-50 mb-2" href="#"><i class="fa fa-angle-right mr-2"></i>Guides</a>
                    <a class="text-white-50 mb-2" href="#"><i class="fa fa-angle-right mr-2"></i>Testimonial</a> -->
                    <!-- <a class="text-white-50" href="#"><i class="fa fa-angle-right mr-2"></i>Blog</a> -->
                </div>
            </div>
            <div class="col-lg-3 col-md-6 mb-5">
                <h5 class="text-white text-uppercase mb-4" style="letter-spacing: 5px;">Colaboradores</h5>
                <div class="d-flex flex-column justify-content-start">
                    <a class="text-white-50 mb-2" href="https://web.facebook.com/revoluciontics"><i class="fa fa-angle-right mr-2"></i>REVOLUCION TIC S.A.S</a>
                    <!-- <a class="text-white-50 mb-2" href="#"><i class="fa fa-angle-right mr-2"></i>Destination</a>
                    <a class="text-white-50 mb-2" href="#"><i class="fa fa-angle-right mr-2"></i>Services</a>
                    <a class="text-white-50 mb-2" href="#"><i class="fa fa-angle-right mr-2"></i>Packages</a>
                    <a class="text-white-50 mb-2" href="#"><i class="fa fa-angle-right mr-2"></i>Guides</a>
                    <a class="text-white-50 mb-2" href="#"><i class="fa fa-angle-right mr-2"></i>Testimonial</a>
                    <a class="text-white-50" href="#"><i class="fa fa-angle-right mr-2"></i>Blog</a> -->
                </div>
            </div>
            <div class="col-lg-3 col-md-6 mb-5">
                <h5 class="text-white text-uppercase mb-4" style="letter-spacing: 5px;">Contacto</h5>
                <p><i class="fa fa-map-marker-alt mr-2"></i>Monteria, Córdoba, Colombia</p>
                <p><i class="fa fa-phone-alt mr-2"></i>+57 312 5637781</p>
                <p><i class="fa fa-envelope mr-2"></i>costabrisaweb@gmail.com</p>
                <div class="w-100">
                </div>
            </div>
        </div>
    </div>
    <div class="container-fluid bg-dark text-white border-top py-4 px-sm-3 px-md-5" style="margin-top: -90px;border-color: rgba(256, 256, 256, .1) !important;">
        <div class="row">
            <div class="col-lg-6 text-center text-md-left mb-3 mb-md-0">
               
            </div>
            <div class="col-lg-6 text-center text-md-right" style="margin-top: 10px;">
                <p class="m-0 text-white-50">Designed by <a href="#">CostaBrisa</a>
                </p>
            </div>
        </div>
    </div>
    <!-- Footer End -->
    

   <!-- eslint-enable -->
</template>
<script>

import { ref, watchEffect } from 'vue'
import axios from 'axios'
export default {
  data() {
    return {
      notificaciones: [],
    };
  },
  setup() {
    // Declara las variables reactivas usando ref()
    const usuarioAutenticado = ref(localStorage.getItem('usuarioAutenticado'));
    const usuarioRol = ref(localStorage.getItem('usuarioRol'));
    const imagenUsuario = ref(localStorage.getItem('imagenUsuario'));
    const nombreUsuario = ref(localStorage.getItem('nombreUsuario'));
    // Observa los cambios en las variables reactivas
    watchEffect(() => {
      usuarioAutenticado.value = localStorage.getItem('usuarioAutenticado');
      usuarioRol.value = localStorage.getItem('usuarioRol');
      imagenUsuario.value = localStorage.getItem('imagenUsuario');
      nombreUsuario.value = localStorage.getItem('nombreUsuario');
    });
    const ocultarMenu = () => {
      // Ocultar el menú al presionar el botón de "Iniciar sesión"
      const navbarCollapse = document.getElementById('navbarCollapse');
      if (navbarCollapse.classList.contains('show')) {
        navbarCollapse.classList.remove('show');
      }
    };

    return {
      usuarioAutenticado,
      usuarioRol,
      imagenUsuario,
      nombreUsuario,
      ocultarMenu
    };
  },
  created(){
   
  this.getNotificacion();
  console.log(this.notificaciones)
  },
  
  
  methods: {
    getNotificacion(){
  let idc= localStorage.getItem('customerId');
    axios.get(`${process.env.API}/pedirServicioc/${idc}`)
      .then((response) => {
        this.notificaciones = response.data.filter(e=> e.revision === true)
        
      })
      .catch((error) => {
        console.log(error);
      });
},
    cerrarSesion() {
      // Limpiar la información de autenticación y del usuario
      localStorage.clear();

      // Otra lógica de cierre de sesión si es necesario
      // Actualizar el estado local de autenticación y del usuario
      this.usuarioAutenticado = false;
      this.usuarioRol = '';
      this.imagenUsuario = '';
      this.nombreUsuario = '';
      this.logout()
    },
    logout() {
        this.$auth0.logout({ logoutParams: { returnTo: window.location.origin } });
      }
  },
       
 
};
</script>


<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 5px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}

/* Estilos personalizados para el navbar */
body {
  padding-top: 10px; /* Ajusta el cuerpo de la página para evitar que el navbar oculte el contenido */
}

.navbar {
  background-color: #3498db; /* Color de fondo del navbar */
  height: 90px;
  width: 100%;
}

.navbar-brand,
.navbar-nav .nav-link {
  color: #fff; /* Color del texto del navbar */
}

.navbar-toggler-icon {
  background-color: rgba(255, 255, 255, 0.142)ccc; /* Color del ícono del botón de alternar */
}

/* Animación personalizada para el navbar */
@keyframes slideInLeftToRight {
  from {
    opacity: 0;
    transform: translate3d(-100%, 0, 0);
  }
  to {
    opacity: 1;
    transform: none;
  }
}

/* Estilos para dispositivos móviles */
@media (max-width: 767px) {
  .navbar-collapse {
    animation: slideInLeftToRight 0.5s ease;
    background-color: rgb(255, 255, 255); /* Fondo blanco para el menú desplegable en dispositivos móviles */
  }
}
.usuario-info {
  display: flex;
  align-items: center;
}

.avatar {
  width: 50px; /* ajusta el tamaño según tus preferencias */
  height: 50px; /* ajusta el tamaño según tus preferencias */
  border-radius: 50%; /* para hacerlo circular, ajusta según tus preferencias */
  margin-right: 10px; /* ajusta el margen según tus preferencias */
}

.nombre-usuario {
  margin: 0; /* asegúrate de quitar cualquier margen predeterminado si es necesario */
}
@keyframes glowAnimation {
            0% {
                box-shadow: 0 0 5px rgba(255, 255, 255, 0.5);
            }
            50% {
                box-shadow: 0 0 20px rgba(255, 255, 255, 0.7);
            }
            100% {
                box-shadow: 0 0 5px rgba(255, 255, 255, 0.5);
            }
        }

        .glow-text {
            display: inline-block;
            animation: glowAnimation 2s infinite;
        }

        h1 {
            font-family: 'Arial', sans-serif;
            font-size: 2em;
            margin: 0;
        }

        .text-primary {
            color: #007bff;
        }

        .text-dark {
            color: #333;
        }



/* Estilos generales para el enlace */
.navbar-brand {
  text-decoration: none;
  color: #000; /* Color del texto sin pasar el cursor */
  position: relative;
  overflow: hidden;
  transition: color 0.3s ease;
}

/* Efecto de sombreado y transición al pasar el cursor */
.navbar-brand::before {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 2px;
  background-color: #007bff; /* Color del sombreado al pasar el cursor */
  transform: scaleX(0);
  transform-origin: bottom right;
  transition: transform 0.3s ease;
}

.navbar-brand:hover {
  color: #007bff; /* Color del texto al pasar el cursor */
}

.navbar-brand:hover::before {
  transform: scaleX(1);
  transform-origin: bottom left;
}




</style>
