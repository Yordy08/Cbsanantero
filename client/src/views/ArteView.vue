<template>
  <div class="container">
    <hr>
    <div class="row justify-content-center">
      <div class="col-lg-4 col-md-6 col-sm-12" v-for="artesania in artesanias" :key="artesania.id">
        <div class="card mb-4">
          <img @click="enviarIdGaleria(artesania._id, artesania.customer_id)" :src="artesania.image" class="card-img-top" alt="Artesanía" />
          <div class="card-body">
            <h5 class="card-title">{{ artesania.name }}</h5>
            <p class="card-text">{{ artesania.description }}</p>
            <p class="card-text"><strong>Dirección:</strong> {{ artesania.address }}</p>
            <p class="card-text"><strong>Teléfono:</strong> {{ artesania.phone }}</p>
            <div v-if="customerRol === 'Cliente'">
              <button class="btn btn-primary" data-toggle="modal" data-target="#solicitudModal" @click="enviarcid(artesania.customer_id, artesania._id)">Obtener Servicio</button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Modal -->
    <div class="modal" id="solicitudModal" tabindex="-1" role="dialog" aria-labelledby="solicitudModalLabel" aria-hidden="true">
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <!-- Modal Header -->
          <div class="modal-header">
            <h4 class="modal-title" id="solicitudModalLabel">Solicitar Servicio</h4>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="closeModal"></button>
          </div>
          <!-- Modal Body -->
          <div class="modal-body">
            <b><p>Creando servicio:</p></b>
            <!-- Formulario con campos solicitados -->
            <form>
              <div class="form-group">
                <label for="nombre">Nombres:</label>
                <input type="text" class="form-control" id="nombre" placeholder="Describa nombre y apellido" v-model="formData.nombre">
              </div>
              <div class="form-group">
                <label for="correo">Correo:</label>
                <input type="email" class="form-control" placeholder="Escriba su correo" id="apellidos" v-model="formData.correo">
              </div>
              <div class="form-group">
                <label for="celular">Numero de celular:</label>
                <input type="text" class="form-control" id="celular" placeholder="eje +57 304 000 4445" v-model="formData.celular">
              </div>
              <div class="form-group">
                <label for="description">Descripcion:</label>
                <textarea id="description" placeholder="Deescriba su solicitud" rows="3"></textarea>
              </div>
              <button type="button" class="btn btn-success" @click="submitForm">Solicitar</button>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>


<script>
import axios from 'axios';

export default {
  data() {
    return {
      artesanias: [],
      customerRol: '',
      userRole: 'usuario', // Asume que userRole es 'usuario' por defecto. Adáptalo según la información real en tu aplicación.
      formData: {
        _id:null,
        nombre: '',
/*         apellidos: '', */
        /* descripcion: '',   esta descripcion no esta llegando debido a que utilizo un area de texto y identifica el ID  */
        celular: '',
        correo: '',
        customerId: null,
      },
    };
  },
  mounted() {
    this.customerRol = localStorage.getItem('usuarioRol');
    this.fetchArtesanias();
  },
  methods: {
    async fetchArtesanias() {
      try {
        const response = await axios.get(`${process.env.API}/Artesania`);
        this.artesanias = response.data;
        console.log(this.artesanias)
      } catch (error) {
        console.error('Error al obtener las artesanías', error);
      }
    },
    enviarIdGaleria(id,cid) {
      localStorage.setItem('negocioId',id);
      localStorage.setItem('customerNegocioId', cid);
      this.$router.push({ path: '/artegaleria' });
    },

    enviarcid(cid,id){
      this.formData.customerId=cid;
      this.formData._id=id;
      console.log(this.formData)
    },
    openModal(customerId) {
      this.selectedCustomerId = customerId;
      this.formData.customerId = customerId;
      this.modalOpen = true;
    },
    closeModal() {
      this.selectedCustomerId = null;
      this.formData = {
        _id:null,
        nombre: '',
        apellidos: '',
        descripcion: '',
        celular: '',
        correo: '',
        customerId: null,
      };
    },
    submitForm() {
      console.log('CustomerID seleccionado:', this.selectedCustomerId);
      console.log('Formulario enviado con datos:', this.formData);
      this.closeModal();
      // Puedes agregar aquí la lógica para enviar el formulario al servidor si es necesario.
    },
  },
};
</script>

<style scoped>
.card {
  transition: transform 0.3s ease-in-out, box-shadow 0.3s ease-in-out;
}

.card:hover {
  transform: scale(1.05);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  border-color: blue; /* Color del borde al pasar el cursor */
}

.card-img-top {
  width:auto; /* Ancho deseado */
  height:400px; /* Altura deseada */
  object-fit: cover; /* Para asegurar que la imagen se ajuste correctamente sin distorsión */
}
</style>
