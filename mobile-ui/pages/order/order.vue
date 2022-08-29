<template>
	<view>
		<u-navbar back-text="我的订单" :borderBottom="false" :background="background" :back-text-style="{'color':'#fff'}"
			backIconColor="#fff"></u-navbar>
			  <u-subsection :list="list" :current="curNow" @change="sectionChange"></u-subsection>
			  <view class="" v-if="foodPublishList === null ||foodPublishList.length===0 " style="margin-top: 50%;">
			  	<u-empty
			  	        mode="data"
			  	        icon="https://cdn.uviewui.com/uview/demo/empty/data.png"
			  	>
			  	</u-empty>
			  </view>
			  		
		
		<view>
			<view style="display: flex;flex-direction: row;align-items: center;margin-left: 16upx;margin-top: 20upx;"  v-for="(item,index) in foodPublishList" :key="index">
				<img :src="item.image" mode="aspectFill" style="width: 200upx;height: 200upx;border-radius: 20upx;"></img>
				<view style="display: flex;flex-direction: column;margin-left: 10upx;">
					<view class="" style="display: flex;flex-direction: row;align-items: center;justify-content: space-between;width: 500upx;margin-bottom: 20upx;">
						<view style="font-size: 16px;font-weight: bold;">{{item.name}}</view>
						
						<view v-if="item.status===1">待付款</view>
						<view v-else-if="item.status===2">已付款</view>
						<view v-else>已完成</view>
						
					</view>	
					<view style="display: flex;flex-direction: row;align-items: center;justify-content: space-between;">
						<view style="font-size: 12px;">订单编号:{{item.orderNo}}</view>
						<view style="font-size: 18px;font-weight: bold;color: red;">￥{{item.price}}</view>
					</view>
					
					<view>{{item.description}}</view>
					<view style="display: flex;flex-direction: row;align-items: center;justify-content: space-between;">
						<view class="" style="font-size: 12px;margin-top: 10upx;">创建时间:{{item.createTime}}</view>
						<view v-show="item.status===1" @click="updateOrderStatus(item.id,1)">
							<view  style="width: 120upx;height: 50upx;font-size: 20upx;background: #19BE6B;
							border-radius: 50upx;text-align: center;line-height: 50upx;color: white;">去支付</view>
						</view>
						<view v-show="item.status===2" @click="updateOrderStatus(item.id,2)">
							<view  style="width: 120upx;height: 50upx;font-size: 20upx;background: #19BE6B;
							border-radius: 50upx;text-align: center;line-height: 50upx;color: white;">确认订单</view>
						</view>
					</view>
					
				</view>
				
			</view>
			
			
			
			
		</view>
		
		
		
	</view>
</template>

<script>
	import {mine_collect_list,food_order_list,food_order_update_status} from '../../api/food.js'
	
	export default {
		data() {
			return {
				background: {
					'background-image': 'linear-gradient(45deg, rgb(28, 187, 180), rgb(141, 198, 63))'
				},
				list: ['待付款','已付款','已完成'],
				curNow: 0,
				userinfo:{},
				foodPublishList:[]
			}
		},
		onLoad() {
			this.sectionChange(0)
		},
		methods: {
			
			updateOrderStatus(id,type){
				if(type===2){
					let params={}
					params.id=id
					params.status=type
					
					food_order_update_status(params).then(res=>{
						uni.showToast({
							title:'确认成功'
						})
						this.curNow=2;
						this.sectionChange(2)
					})
				}else{
					let params={}
					params.id=id
					params.status=type
					food_order_update_status(params).then(res=>{
						uni.showToast({
							title:'支付成功'
						})
						this.curNow=1;
						this.sectionChange(1)
					})
				}
				
			},
			sectionChange(index) {
					this.curNow = index;
					console.log(index)
					if(index===0){
						let params={}
						params.type=1
						food_order_list(params).then(res=>{
							console.log(res)
							this.foodPublishList=res.data
						})
					}else if(index===1){
						let params={}
						params.type=2
						food_order_list(params).then(res=>{
							console.log(res)
							this.foodPublishList=res.data
						})
					}else{
						let params={}
						params.type=3
						food_order_list(params).then(res=>{
							console.log(res)
							this.foodPublishList=res.data
						})
					}
					
							
			}

		}
	}
</script>

<style lang="scss" scoped>
</style>
