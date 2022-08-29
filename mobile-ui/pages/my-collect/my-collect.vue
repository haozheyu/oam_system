<template>
	<view>
		<u-navbar back-text="我的收藏" :borderBottom="false" :background="background" :back-text-style="{'color':'#fff'}"
			backIconColor="#fff"></u-navbar>
			  <u-subsection :list="list" :current="curNow" @change="sectionChange"></u-subsection>
			  <view class="" v-if="foodPublishList === null" style="margin-top: 50%;">
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
					<view class="" style="display: flex;flex-direction: row;align-items: center;justify-content: space-between;width: 500upx;margin-bottom: 40upx;">
						<view style="font-size: 16px;font-weight: bold;">{{item.name}}</view>
						
					</view>	
					<view>{{item.description}}</view>
					<view class="" style="font-size: 12px;margin-top: 10upx;">创建时间:{{item.createTime}}</view>
				</view>
				
			</view>
			
			
			
			
		</view>
		
		
		
	</view>
</template>

<script>
	import {mine_collect_list} from '../../api/food.js'
	
	export default {
		data() {
			return {
				background: {
					'background-image': 'linear-gradient(45deg, rgb(28, 187, 180), rgb(141, 198, 63))'
				},
				list: ['菜品'],
				curNow: 0,
				userinfo:{},
				foodPublishList:[]
			}
		},
		onLoad() {
			this.sectionChange(0)
		},
		methods: {
			sectionChange(index) {
					this.curNow = index;
					console.log(index)
					if(index===0){
						mine_collect_list({}).then(res=>{
							console.log(res)
							this.foodPublishList=res.data
						})
					}else{

						mine_collect_list({}).then(res=>{
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
