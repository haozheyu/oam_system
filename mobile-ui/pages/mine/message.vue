<template>
	<view class="page">
	<u-navbar title="消息" backIconColor="#fff" titleColor="#fff" :background="background"></u-navbar>
	<view class="home">
		<view v-for="(item,index) in notifyList" :key="index" class="u-p-b-20">
			<view class="time">{{item.createTime}}</view>
			<view class="box">
				<view class="box-title">注册成功通知！</view>
				<view class="box-content u-tips-color" :v-html="item.content">{{item.content}}</view>
			</view>
		</view>
	</view>
	</view>
</template>

<script>
	import {notify_list} from '@/api/notify.js'
	
	export default {
		data() {
			return {
				notifyList:[],
				background: {
					'background-image': 'linear-gradient(45deg, rgb(28, 187, 180), rgb(141, 198, 63))'
				},
				content:'成功注册脍炙食谱小程序，您的加入就是我前进的动力~'
			}
		},
		onLoad() {
			this.getNotifyList()
		},
		methods:{
			
			getNotifyList(){
				notify_list({}).then(res=>{
					console.log(res)
					this.notifyList=res.data
				})
			}
			
		}
	}
</script>

<style lang="scss">
	page{
		background-color: #ededed;
	}
	.time{
		text-align: center;
		color: #39CCCC;
		padding: 30rpx 0 20rpx 0;
	}
	.box{
		padding: 30rpx;
		background-color: #FFFFFF;
		border-radius: 6rpx;
		.box-title{
			text-align: left;
			font-size: 32rpx;
			font-weight: 700;
		}
		.box-content{
			margin-top: 10rpx;
			font-size: 28rpx;
		}
	}
</style>
