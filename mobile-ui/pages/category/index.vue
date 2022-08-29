<template>
	<view class="">
		<u-navbar title="分类" :is-back="false" :background="background" titleColor="#fff"></u-navbar>
		<view class="u-wrap">
			<view class="u-search-box" @click="toSearch">
				<view class="u-search-inner">
					<u-icon name="search" color="#909399" :size="28"></u-icon>
					<text class="u-search-text">请输入关键词搜索</text>
				</view>
			</view>
			<view class="u-menu-wrap">
				<scroll-view scroll-y scroll-with-animation class="u-tab-view menu-scroll-view" :scroll-top="scrollTop">
					<view v-for="(item,index) in categoryList" :key="index" class="u-tab-item" :class="[current==index ? 'u-tab-item-active' : '']"
					 :data-current="index" @tap.stop="swichMenu(index,item)">
						<text class="u-line-1">{{item.name}}</text>
					</view>
				</scroll-view>
				<block v-for="(item,index) in categoryList" :key="index">
					<scroll-view scroll-y class="right-box" v-if="current==index">
						<view class="page-view">
							<view class="class-item">
								<view class="item-title">
									<text>{{item.name}}</text>
								</view>
								<view class="item-container">
									<view class="thumb-box" v-for="(item1, index1) in categoryDetailList" :key="index1" @click="toDetail(item1.id)">
										<image class="item-menu-image" :src="item1.imgUrl"></image>
										<view class="item-menu-name">{{item1.name}}</view>
									</view>
								</view>
							</view>
						</view>
					</scroll-view>
				</block>
			</view>
		</view>
	</view>
</template>


<script>
	import {food_category_list} from '../../api/food_category.js'
	import {food_categroy_detail,get_windows_by_restaurant} from '../../api/food.js'
	export default {
		data() {
			return {
				background: {
					'background-image': 'linear-gradient(45deg, rgb(28, 187, 180), rgb(141, 198, 63))'
				},
				img:'/static/pic.jpg',
			
				scrollTop: 0, 
				oldScrollTop: 0,
				current: 0, 
				menuHeight: 0, 
				menuItemHeight: 0, 
				itemId: '', 
				tabbar: [],
				menuItemPos: [],
				arr: [],
				scrollRightTop: 0, 
				timer: null, 
				categoryList:[],
				categoryDetailList:[]
			}
		},
		onShow() {
			let serviceItemId = uni.getStorageSync('serviceItemId')
			console.log("1111111111111===>"+serviceItemId)
			this.getCategroyWindowsList(serviceItemId)
			// this.getCategroyList()
			let parmas={}
			parmas.categoryId=8
			
			food_categroy_detail(parmas).then(res=>{
				console.log(res)
				this.categoryDetailList=res.data
			})
		},
		onLoad(e) {
			// let serviceItemId = uni.getStorageSync('serviceItemId')
			// console.log("1111111111111===>"+serviceItemId)
			// this.getCategroyWindowsList(serviceItemId)
			// // this.getCategroyList()
			// let parmas={}
			// parmas.categoryId=8
			
			// food_categroy_detail(parmas).then(res=>{
			// 	console.log(res)
			// 	this.categoryDetailList=res.data
			// })
			
		},
		methods: {
			getCategroyWindowsList(restaurantId){
				if(restaurantId){
					let params={}
					params.restaurantId=restaurantId
					get_windows_by_restaurant(params).then(res=>{
						this.categoryList=res.data
					})
				}else{
					let params={}
					params.restaurantId=1
					get_windows_by_restaurant(params).then(res=>{
						this.categoryList=res.data
					})
				}
				
				
			},
			
			getCategroyList(){
				food_category_list({}).then(res=>{
					console.log(res)
					this.categoryList=res.data
				})
			},
			
			toSearch(){
				this.$common.navigateTo('/pages/index/search')
			},
			toDetail(id){
				console.log(id)
				// this.$common.navigateTo('/pages/index/detail')
				uni.navigateTo({
					url:'/pages/index/detail?id='+id
				})
			},
			// 点击左边的栏目切换
			async swichMenu(index,item) {
				if(index == this.current) return ;
				this.current = index;
				console.log(item.id)
				let parmas={}
				parmas.categoryId=item.windowsId
				food_categroy_detail(parmas).then(res=>{
					console.log(res)
					this.categoryDetailList=res.data
				})
				// 如果为0，意味着尚未初始化
				if(this.menuHeight == 0 || this.menuItemHeight == 0) {
					await this.getElRect('menu-scroll-view', 'menuHeight');
					await this.getElRect('u-tab-item', 'menuItemHeight');
				}
				
				// 将菜单菜单活动item垂直居中
				this.scrollTop = index * this.menuItemHeight + this.menuItemHeight / 2 - this.menuHeight / 2;
			},
			// 获取一个目标元素的高度
			getElRect(elClass, dataVal) {
				new Promise((resolve, reject) => {
					const query = uni.createSelectorQuery().in(this);
					query.select('.' + elClass).fields({size: true},res => {
						// 如果节点尚未生成，res值为null，循环调用执行
						if(!res) {
							setTimeout(() => {
								this.getElRect(elClass);
							}, 10);
							return ;
						}
						this[dataVal] = res.height;
					}).exec();
				})
			}
		}
	}
</script>

<style lang="scss" scoped>
	.u-wrap {
		height: calc(100vh);
		/* #ifdef H5 */
		height: calc(100vh - var(--window-top));
		/* #endif */
		display: flex;
		flex-direction: column;
	}
	
	.u-search-box {
		padding: 18rpx 30rpx;
	}
	
	.u-menu-wrap {
		flex: 1;
		display: flex;
		overflow: hidden;
	}
	
	.u-search-inner {
		background-color: rgb(234, 234, 234);
		border-radius: 100rpx;
		display: flex;
		justify-content: center;
		align-items: center;
		padding: 16rpx 16rpx;
	}
	
	.u-search-text {
		font-size: 26rpx;
		color: $u-tips-color;
		margin-left: 10rpx;
	}
	
	.u-tab-view {
		width: 200rpx;
		height: 100%;
	}
	
	.u-tab-item {
		height: 110rpx;
		background: #f6f6f6;
		box-sizing: border-box;
		display: flex;
		align-items: center;
		justify-content: center;
		font-size: 26rpx;
		color: #444;
		font-weight: 400;
		line-height: 1;
	}
	
	.u-tab-item-active {
		position: relative;
		color: #6BC362;
		font-size: 30rpx;
		font-weight: 600;
		background: #fff;
	}
	
	.u-tab-item-active::before {
		content: "";
		position: absolute;
		border-left: 4px solid #6BC362;
		height: 32rpx;
		left: 0;
		top: 39rpx;
	}
	
	.u-tab-view {
		height: 100%;
	}
	
	.right-box {
		background-color: rgb(250, 250, 250);
	}
	
	.page-view {
		padding: 16rpx;
	}
	
	.class-item {
		margin-bottom: 30rpx;
		background-color: #fff;
		padding: 16rpx;
		border-radius: 8rpx;
	}
	
	.class-item:last-child {
		min-height: 100vh;
	}
	
	.item-title {
		font-size: 26rpx;
		color: $u-main-color;
		font-weight: bold;
	}
	
	.item-menu-name {
		font-weight: normal;
		font-size: 24rpx;
		color: $u-main-color;
	}
	
	.item-container {
		display: flex;
		flex-wrap: wrap;
	}
	
	.thumb-box {
		width: 33.333333%;
		display: flex;
		align-items: center;
		justify-content: center;
		flex-direction: column;
		margin-top: 20rpx;
	}
	
	.item-menu-image {
		width: 150rpx;
		height: 150rpx;
		border-radius: 6rpx;
	}
</style>
