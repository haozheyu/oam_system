<template>
	<view>
		<u-navbar back-text="登录" :customBack="back" :borderBottom="false" :background="background"
			:back-text-style="{'color':'#fff'}" backIconColor="#fff"></u-navbar>
		<view class="home center">
			<view>
				<!-- logo -->
				<u-image src="/static/banner3.jpg" height="450" borderRadius="12"></u-image>
			</view>
			<view class="u-p-t-30">
				<view class="wxname">欢迎登录食堂点餐小程序</view>
				<view class="u-font-12 u-tips-color u-m-b-50">欢迎登录食堂点餐小程序</view>
				<view style="width: 90%;margin-left: 50%;transform: translate(-50%);margin-top: 20upx;">
					<u-input placeholder="请输入手机号"  v-model="loginForm.mobile" shape="circle" clearable></u-input>
				</view>
				
				<view style="width: 90%;margin-left: 50%;transform: translate(-50%);margin-top: 20upx;">
					<u-input placeholder="请输入密码"  type="password" v-model="loginForm.password" shape="circle" clearable></u-input>
				</view>
				<view style="margin-top: 40upx;">
					
					<u-button type="success" shape="circle" @click="userLogin()" >登录</u-button>
				</view>
				
				<view class="center u-tips-color u-p-t-50" @click="toRegister()">注册</view>
			</view>
		</view>
	</view>
</template>

<script>
	import {login} from '../../api/user.js'

	import {
		set
	} from '../../utils/local.js'
	export default {
		data() {
			return {
				background: {
					'background-image': 'linear-gradient(45deg, rgb(28, 187, 180), rgb(141, 198, 63))'
				},
				loginForm:{}
			}
		},
		onLoad() {

		},
		methods: {
			back() {
				return this.$common.switchTabTo('/pages/index/index')
			},
			userLogin() {
				// 登录方法
				if(this.loginForm.mobile ==='' || this.loginForm.mobile===null){
					uni.showToast({
						icon:'none',
						title:'请输入手机号'
					})
					return
				}
				if(this.loginForm.password ==='' || this.loginForm.password===null){
					uni.showToast({
						icon:'none',
						title:'请输入密码'
					})
					return
				}
				login(this.loginForm).then(res=>{
					if(res.status===200){
						console.log(res)
						uni.showToast({
							title:'登录成功',
							duration:2000
						})
						set('userId',res.data.userId)
						set('avatar',res.data.avatar)
						set('nickname',res.data.nickname)
						set('token',res.data.token)
						
						uni.switchTab({
							url:'../mine/index'
						})
						
						
					}
				})
			},
			
			toRegister(){
				uni.navigateTo({
					url:'../register/register'
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
