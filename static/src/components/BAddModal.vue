<template>
  <MDBModal tabindex="-1" v-model="show">
    <form @submit.prevent="save" @keypress="formKeyPress">
      <MDBModalHeader>
        <MDBModalTitle>{{ caption }}</MDBModalTitle>
      </MDBModalHeader>

      <MDBModalBody>
        <slot/>
        <div class="alert alert-danger mt-3" v-if="error">{{ error }}</div>
      </MDBModalBody>

      <MDBModalFooter>
        <MDBBtn color="primary" type="submit" :disabled="saving">{{ addBtn }}</MDBBtn>
        <MDBBtn color="secondary" :disabled="saving" @click="show=false">{{ cancelBtn }}</MDBBtn>
      </MDBModalFooter>
    </form>
  </MDBModal>
</template>

<script lang="ts">
import {defineComponent, PropType} from 'vue';
import {MDBBtn, MDBModal, MDBModalBody, MDBModalFooter, MDBModalHeader, MDBModalTitle} from "mdb-vue-ui-kit";

export default defineComponent({
  name: 'BAddModal',
  components: {
    MDBBtn, MDBModal, MDBModalBody, MDBModalFooter, MDBModalHeader, MDBModalTitle
  },
  props: {
    caption: String,
    addBtn: {
      type: String,
      default: 'Add'
    },
    cancelBtn: {
      type: String,
      default: 'Cancel'
    },
    modelValue: Boolean,
    callback: {
      // eslint-disable-next-line
      type: Function as PropType<() => Promise<any>>,
      required: true,
    }
  },
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
      this.error = ''
    }
  },
  methods: {
    save() {
      this.saving = true
      this.error = ''

      this.callback().then(() => {
        this.show = false
      }).catch((err) => {
        this.error = err
      }).finally(() => {
        this.saving = false
      })
    },

    formKeyPress(e: KeyboardEvent) {
      if (e.keyCode === 13 && (e.target as HTMLElement).nodeName != "TEXTAREA") {
        e.preventDefault();
        this.save()
      }
    }

  }
});
</script>