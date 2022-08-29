<template>
	<view>
		<u-navbar back-text="详情" :borderBottom="false" :background="background" :back-text-style="{'color':'#fff'}"
			backIconColor="#fff"></u-navbar>
		<view>
			<u-image :src="detail.imgUrl" height="550" borderRadius="0"></u-image>
			<view class="home">
				<view class="u-flex u-row-between">
					<view class="title-blod">{{detail.name}}</view>
					
					<view>
						<u-icon name="heart" color="#6BC362" size="40" label="收藏" v-if="!detail.isCollect" @click="collect(1)"></u-icon>
						<u-icon name="heart-fill" color="#6BC362" size="40" label="已收藏" v-else @click="collect(2)"></u-icon>
					</view>
				</view>
				<view class="u-m-t-20">
				<!-- 	<u-tag v-for="(item,index) in 1" :key="index" :text="detail.restaurantName" mode="dark" class="tag"
						type="success"></u-tag> -->
					
					<u-tag v-for="(item,index) in 1" :key="index" :text="detail.windowsName" mode="dark" class="tag"
						type="success"></u-tag>
				</view>
			<!-- 	<view class="u-m-t-30 u-flex u-row-between">
					<view class="times">烹饪时间：{{detail.cookTime}}分钟</view>
					<!-- <view class="times">用餐人数：2人</view> -->
				<!-- </view> -->
				<view class="u-m-t-30" style="margin-left: 20upx;">
					<view class="title-blod u-m-b-10">菜品描述</view>
					<view>{{detail.description}}</view>
				</view>
				<u-gap height="30"></u-gap>
				
				<view class="u-m-t-30" style="margin-left: 20upx;">
					<view class="title-blod">步骤</view>
					<view class="" v-if="detail.foodMakeWays">
						<view v-for="(item,index) in detail.foodMakeWays" :key="index" class="u-m-t-20">
							<view class="flex-process">
								<view>{{index+1}}.{{item}}</view>
							</view>
							<!-- <u-image src="/static/pic.jpg" height="350" border-radius="0"></u-image> -->
						</view>
					</view>	
					<view class="" v-else>
						暂无食材制作步骤哦~
					</view>
					
				</view>
		
			</view>
		</view>
		
		<view style="position: fixed;bottom: 0upx;margin-bottom: 68upx;padding: 0 30upx;" @click="createOrder">
			<view style="width: 690upx;height: 80upx;background: #6BC362;border-radius: 40upx;
			text-align: center;line-height: 80upx;color: white;">立即下单</view>
		</view>
		
		
		<view class="popup" v-show="show">
			<view class="popup-info">
				<view style="display: flex;flex-direction: column;align-items: center;">
					<view style="color: #191919;font-size: 32upx;margin-top: 48upx;">是否确认下单该菜品吗？</view>
				</view>
				
				<view style="display: flex;flex-direction: row;align-items: center;justify-content: space-between;margin-top: 56upx;">
					<view style="background: #F4F4F4;color: #191919;font-size: 28upx;width: 240upx;height: 80upx;margin-left: 30upx;
					border-radius: 32upx;text-align: center;line-height: 80upx;" @tap="cancel">取消</view>
					<view style="background: #6BC362;color: #FFFFFF;font-size: 28upx;width: 240upx;
					height: 80upx;margin-right: 30upx;border-radius: 32upx;text-align: center;line-height: 80upx;" @tap="confrim">确认</view>
				</view>
			</view>
		</view>
		
		<!-- top s -->
		<u-back-top :scrollTop="scrollTop" mode="circle" bottom="200" right="40" top="600" :icon="icon"
			:icon-style="iconStyle" :tips="tips" :custom-style="customStyle">
		</u-back-top>
	</view>
</template>

<script>
	import {food_detail,food_collect,food_order_create} from '../../api/food.js'
	
	export default {
		data() {
			return {
				background: {
					'background-image': 'linear-gradient(45deg, rgb(28, 187, 180), rgb(141, 198, 63))'
				},
				detail: {},
				show: false,
				scrollTop: 0,
				icon: 'arrow-up',
				iconStyle: {
					color: '#ffffff',
					fontSize: '30rpx'
				},
				tips: '顶部',
				customStyle: {
					backgroundColor: '#6BC362',
					color: '#ffffff'
				},
				
			}
		},
		onLoad(e) {
			console.log(e.id)
		
			this.getFoodDetail(e.id)
		},
		methods: {
			createOrder(){
				this.show=true
				// let params={}
				// params.foodId=this.detail.id;
				// food_order_create(params).then(res=>{
					
				// })
			},
			cancel(){
				let params={}
				params.status=1;
				params.foodId=this.detail.id;
				food_order_create(params).then(res=>{
					
					this.show=false;
					uni.navigateTo({
						url:'../order/order'
					})
				})
			},
			
			confrim(){
				let params={}
				params.status=2;
				params.foodId=this.detail.id;
				food_order_create(params).then(res=>{
					uni.showToast({
						title:'下单成功',
						duration:2000
					})
					this.show=false;
					uni.navigateTo({
						url:'../order/order'
					})
				})
			},
			
			collect(type){
				console.log(type)
				if(type===1){
					this.detail.isCollect=true
					uni.showToast({
						title:'收藏成功'
					})
					let params={}
					params.foodId=this.detail.id
					params.type=1
					food_collect(params).then(res=>{
						console.log(res)
					})
				}else{
					this.detail.isCollect=false
					uni.showToast({
						title:'操作成功'
					})
					let params={}
					params.foodId=this.detail.id
					params.type=2
					food_collect(params).then(res=>{
						console.log(res)
					})
				}
				
			},
			getFoodDetail(id){
				let params={}
				params.id=id
				food_detail(params).then(res=>{
					console.log(res)
					this.detail=res.data
				})
			},
			// top
			onPageScroll(e) {
				this.scrollTop = e.scrollTop;
			},
	
		}
	}
</script>

<style lang="scss" scoped>
	
	.tag {
		margin: 0 20rpx 20rpx 0;
		display: inline-block;
	}

	.bgFA {
		background: #FFAA3E;
	}

	.texts {
		margin-left: 20px;
		width: 40rpx;
		height: 40rpx;
		color: #FFFFFF;
		border-radius: 24rpx;
		font-size: 26rpx;
		padding: 0 20rpx;
	}

	.textp {
		color: #FFFFFF;
		border-radius: 55rpx;
		font-size: 26rpx;
		margin-right: 10rpx;
		padding: 0 10rpx;
		background-image: linear-gradient(45deg, rgb(28, 187, 180), rgb(141, 198, 63));
	}

	.times {
		width: 49%;
		padding: 20rpx 0;
		background-image: linear-gradient(45deg, rgb(28, 187, 180), rgb(141, 198, 63));
		text-align: center;
		color: #FFFFFF;
		font-size: 24rpx;
	}

	.num {
		z-index: 10;
		left: 14rpx;
		font-size: 24rpx;
	}

	.flex-process {
		display: flex;
	}
	.popup-info {
		position: fixed;
		width: 600upx;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		padding: 40upx;
		border-radius: 20upx;
		background-color: #fff;
		z-index: 9999;
	}
	
	.popup {
		position: fixed;
		left: 0;
		right: 0;
		top: 0;
		height: 100vh;
		background-color: rgba(0, 0, 0, 0.6);
		z-index: 9998;
	}
</style>
