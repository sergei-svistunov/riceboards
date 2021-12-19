<template>
  <MDBModal tabindex="-1" v-model="show">
    <MDBModalBody>
      <slot/>
      <div class="alert alert-danger mt-3" v-if="error">{{ error }}</div>
    </MDBModalBody>

    <MDBModalFooter>
      <MDBBtn color="primary" :disabled="saving" @click="doConfirm">{{ yesBtn }}</MDBBtn>
      <MDBBtn color="secondary" :disabled="saving" @click="show=false">{{ noBtn }}</MDBBtn>
    </MDBModalFooter>
  </MDBModal>
</template>

<script lang="ts">
import {defineComponent, PropType} from 'vue';
import {MDBBtn, MDBModal, MDBModalBody, MDBModalFooter} from "mdb-vue-ui-kit";

export default defineComponent({
  name: 'BConfirmModal',
  components: {
    MDBBtn, MDBModal, MDBModalBody, MDBModalFooter
  },
  props: {
    yesBtn: String,
    noBtn: String,
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
    }
  },
  methods: {
    doConfirm() {
      this.saving = true
      this.error = ''

      this.callback().then(() => {
        this.show = false
      }).catch((err) => {
        this.error = err
      }).finally(() => {
        this.saving = false
      })
    }
  }
});
</script>