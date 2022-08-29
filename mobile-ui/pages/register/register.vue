<template>
	<view>
		<u-navbar back-text="注册" :customBack="back" :borderBottom="false" :background="background"
			:back-text-style="{'color':'#fff'}" backIconColor="#fff"></u-navbar>
		<view class="home center">
			<view>
				<!-- logo -->
				<u-image src="/static/banner3.jpg" height="450" borderRadius="12"></u-image>
			</view>
			<view class="u-p-t-30">
				<view class="wxname">欢迎注册广师微食堂</view>
				<view class="u-font-12 u-tips-color u-m-b-50">欢迎注册广师微食堂</view>
				<view style="width: 90%;margin-left: 50%;transform: translate(-50%);margin-top: 20upx;">
					<u-input placeholder="请输入用户昵称" border="surround" v-model="registerForm.nickname" shape="circle" clearable></u-input>
				</view>
				
				<view style="width: 90%;margin-left: 50%;transform: translate(-50%);margin-top: 20upx;">
					<u-input placeholder="请输入手机号" v-model="registerForm.mobile" border="surround" shape="circle" clearable></u-input>
				</view>
				
				<view style="width: 90%;margin-left: 50%;transform: translate(-50%);margin-top: 20upx;">
					<u-input placeholder="请输入密码" border="surround" v-model="registerForm.password" type="password" shape="circle" clearable></u-input>
				</view>
				<view style="width: 90%;margin-left: 50%;transform: translate(-50%);margin-top: 20upx;">
					<u-input placeholder="请再次确认密码" border="surround" type="password" v-model="registerForm.confirmPassword" shape="circle" clearable></u-input>
				</view>
				<view style="margin-top: 40upx;">
					
					<u-button type="success" shape="circle" @click="userRegister()" >注册</u-button>
				</view>
				
				<view class="center u-tips-color u-p-t-50" @click="back()">返回</view>
			</view>
		</view>
	</view>
</template>

<script>
	import {register} from '../../api/user.js'
	export default {
		data() {
			return {
				background: {
					'background-image': 'linear-gradient(45deg, rgb(28, 187, 180), rgb(141, 198, 63))'
				},
				registerForm:{}
			}
		},
		onLoad() {

		},
		methods: {
			back() {
				uni.navigateTo({
					url:'../login/login'
				})
			},
			userRegister() {
				if(this.registerForm.nickname ==='' || this.registerForm.nickname===null){
					uni.showToast({
						icon:'none',
						title:'请输入账号'
					})
					return
				}
				if(this.registerForm.password ==='' || this.registerForm.password===null){
					uni.showToast({
						icon:'none',
						title:'请输入密码'
					})
					return
				}
				if(this.registerForm.password ==='' || this.registerForm.confirmPassword===null){
					uni.showToast({
						icon:'none',
						title:'请再次输入密码'
					})
					return
				}
				if(this.registerForm.password !== this.registerForm.confirmPassword){
					uni.showToast({
						icon:'none',
						title:'两次输入密码不一致'
					})
					return
				}
				// 注册方法
				register(this.registerForm).then(res=>{
					console.log(res)
					if(res.status===200){
						uni.showToast({
							title:'注册成功请登录',
							icon:'success'
						})
					}else{
						uni.showToast({
							title:res.msg,
							icon:'error'
						})
					}
				})
			},

		}
	}
</script>

<style lang="scss">
	page {
		background-color: #fff;
	}

	.center {
		text-align: center;
	}

	.wxname {
		padding-bottom: 20rpx;
		font-size: 38rpx;
		font-weight: bold;
	}
</style>
