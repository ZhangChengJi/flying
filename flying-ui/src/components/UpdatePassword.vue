<template>
  <a-modal
      :title="this.$t('Login.updatePassword')"
      :visible="visible"

      :dialog-style="{ top: '20px' }"
      :maskClosable="false"
      :confirm-loading="confirmLoading"
      @ok="handleOk"
      :okText="this.$t('Button.determine')"
      :cancelText="this.$t('Button.cancel')"
      @cancel="handleCancel"

    >
    <a-form-model ref="ruleForm" :model="ruleForm" :rules="rules" v-bind="layout">
      <a-form-model-item  :label="this.$t('Login.currentPassword')" prop="password">
        <a-input v-model="ruleForm.password"  type="password"/>
      </a-form-model-item>
      <a-form-model-item has-feedback :label="this.$t('Login.newPassword')" prop="pass">
        <a-input v-model="ruleForm.pass" type="password" autocomplete="off" />
      </a-form-model-item>
      <a-form-model-item has-feedback :label="this.$t('Login.confirmPassword')" prop="checkPass">
        <a-input v-model="ruleForm.checkPass" type="password" autocomplete="off" />
      </a-form-model-item>


    </a-form-model>
  </a-modal>

</template>

<script>
import {mapGetters} from "vuex";
import {upPass} from "@/api/user";
import store from "@/store";
import {alertMessage} from "@/utils/alert";

export default {
name: "UpdatePassword",
  data() {
    let checkPending;
    let checkAge = (rule, value, callback) => {
      clearTimeout(checkPending);
      if (!value) {
        return callback(new Error('Please input the age'));
      }
      checkPending = setTimeout(() => {
        if (!Number.isInteger(value)) {
          callback(new Error('Please input digits'));
        } else {
          if (value < 18) {
            callback(new Error('Age must be greater than 18'));
          } else {
            callback();
          }
        }
      }, 1000);
    };

    return {
      visible: false,
      loading: false,
      confirmLoading: false,
      labelCol: { span: 9},
      wrapperCol: { span: 18 },

      form: {
        username: this.username,
        password: '',
        pass: '',
        checkPass: '',
      },

      ruleForm: {
        username: this.username,
        password: '',
        pass: '',
        checkPass: ''
      },
      layout: {
        labelCol: { span: 8 },
        wrapperCol: { span: 14 },
      },
    }
  },
  computed: {
    ...mapGetters(["username"]),
    rules(){
      let validatePass = (rule, value, callback) => {
        if (value === '') {
          callback(new Error(this.$t('Login.pleaseInputNewPassword')));
        } else if(value.length<6){
          callback(new Error(this.$t('Login.lessSixNumber')));
        }
        else {
          if (this.ruleForm.checkPass !== '') {
            this.$refs.ruleForm.validateField('checkPass');
          }
          callback();
        }
      };
      let validatePass2 = (rule, value, callback) => {
        if (value === '') {
          callback(new Error(this.$t('Login.pleaseInputConfirmPassword')));
        } else if (value !== this.ruleForm.pass) {
          callback(new Error(this.$t('Login.twoInputPassword')));
        } else {
          callback();
        }
      };
      return {
        pass: [{ required: true,validator: validatePass, trigger: 'change' }],
        checkPass: [{ required: true, validator: validatePass2, trigger: 'change' }],
        password: [ {required: true, message: this.$t('Login.pleaseInputCurrentPassword'), trigger: 'blur'}],
      }
    }
  },

  methods: {
    showModal() {
      this.visible = true;
    },
    handleOk(e) {
      let _this=this;
      this.$refs.ruleForm.validate(valid => {
        if (valid){
          _this.ruleForm.username=_this.username
          upPass(_this.ruleForm).then(rep =>{
            _this.confirmLoading = false;
            if (alertMessage(_this,rep.data)){
              this.visible = false;
              this.resetForm();
              store.dispatch('FedLogOut').then(() => {
                _this.$router.push({path: '/login'})
              })
            }
          })

        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },
    handleCancel() {
      this.visible = false;
      this.confirmLoading = false;
      this.resetForm();
    },

    resetForm() {
      this.$refs.ruleForm.resetFields();
    },
  }
}
</script>

<style scoped>

</style>