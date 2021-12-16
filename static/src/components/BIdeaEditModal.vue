<template>
  <MDBModal :id="id" tabindex="-1" :labelledby="`${id}Label`" v-model="show">
    <form @submit.prevent="save">
      <MDBModalHeader>
        <MDBModalTitle :id="`${id}Label`">{{ caption }}</MDBModalTitle>
      </MDBModalHeader>

      <MDBModalBody>
        <slot/>
        <div class="alert alert-danger mt-3" v-if="error">{{ error }}</div>
      </MDBModalBody>

      <MDBModalFooter>
        <MDBBtn color="primary" type="submit" :disabled="saving">Save</MDBBtn>
      </MDBModalFooter>
    </form>
  </MDBModal>
</template>

<script lang="ts">
import {defineComponent, PropType} from 'vue';
import {MDBBtn, MDBModal, MDBModalBody, MDBModalFooter, MDBModalHeader, MDBModalTitle} from "mdb-vue-ui-kit";
import api, {IdeasEditReqV1} from "@/api";

export default defineComponent({
  name: 'BIdeaEditModal',
  components: {
    MDBBtn, MDBModal, MDBModalBody, MDBModalFooter, MDBModalHeader, MDBModalTitle
  },
  props: {
    id: String,
    caption: String,
    modelValue: Boolean,
    data: {
      type: Object as PropType<IdeasEditReqV1>,
      required: true
    },
    onSave: {
      type: Function as PropType<() => void>,
      required: true,
    },
  },
  emits: ['update:modelValue'],
  data() {
    return {
      show: this.modelValue,
      saving: false,
      error: '',
    }
  },
  watch: {
    show(value) {
      this.$emit('update:modelValue', value)
    },
    modelValue(value) {
      this.show = value
    }
  },
  methods: {
    save() {
      this.saving = true
      this.error = ''
      api.IdeasEditV1(this.data).then(() => {
        this.onSave()
        this.show = false
      }).catch(err => {
        this.error = err
      }).finally(() => {
        this.saving = false
      })
    }
  }
});
</script>