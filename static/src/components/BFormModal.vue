<template>
  <MDBModal size="lg" :id="id" tabindex="-1" :labelledby="`${id}Label`" v-model="showModal">
    <form @submit.prevent="onSave">
      <MDBModalHeader>
        <MDBModalTitle :id="`${id}Label`">{{ caption }}</MDBModalTitle>
      </MDBModalHeader>

      <MDBModalBody>
        <slot/>
        <div class="alert alert-danger mt-3" v-if="error">{{ error }}</div>
      </MDBModalBody>

      <MDBModalFooter>
        <MDBBtn color="primary" type="submit" :disabled="saving">{{ saveBtn }}</MDBBtn>
      </MDBModalFooter>
    </form>
  </MDBModal>
</template>

<script lang="ts">
import {defineComponent, PropType} from 'vue';
import {MDBBtn, MDBModal, MDBModalBody, MDBModalFooter, MDBModalHeader, MDBModalTitle} from "mdb-vue-ui-kit";

export default defineComponent({
  name: 'BFormModal',
  components: {
    MDBBtn, MDBModal, MDBModalBody, MDBModalFooter, MDBModalHeader, MDBModalTitle
  },
  props: {
    id: String,
    caption: String,
    saveBtn: String,
    show: Boolean,
    onSave: {
      type: Function as PropType<() => void>,
      required: true,
    },
    saving: Boolean,
    error: String,
  },
  data() {
    return {
      showModal: this.show
    }
  },
  watch: {
    show(value) {
      this.showModal = value
    }
  }
});
</script>