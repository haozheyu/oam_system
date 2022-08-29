<template>
	<view class="bg-1">
		<u-waterfall v-model="list" ref="uWaterfall">
			<template v-slot:left="{ leftList }">
				<view class="demo-warter" v-for="(item, index) in leftList" :key="index" @click="toDetail(item.id)">
					<u-lazy-load threshold="50" border-radius="10" :image="item.imgUrl" :index="index">
					</u-lazy-load>
					<view class="demo-title">{{item.name}}</view>

					
					<view class="demo-tag">
						
						<!-- <view class="demo-tag-text">{{item.restaurantName}}</view> -->
						<view class="demo-tag-text">{{item.foodCategoryName}}</view>
						<view style="font-weight: 700;color: red;font-size: 20px;margin-bottom: 8upx;">￥{{item.price}}</view>
					</view>
					<!-- <view class="demo-shop">{{item.description}}</view> -->
					
				</view>
			</template>
			<template v-slot:right="{ rightList }">
				<view class="demo-warter" v-for="(item, index) in rightList" :key="index" @click="toDetail(item.id)">
					<u-lazy-load threshold="50" border-radius="10" :image="item.imgUrl" :index="index">
					</u-lazy-load>
					
		<view class="demo-title">{{item.name}}</view>
		<view class="demo-tag">
			<!-- <view class="demo-tag-text">{{item.restaurantName}}</view> -->
			<view class="demo-tag-text">{{item.foodCategoryName}}</view>
			<view style="font-weight: 700;color: red;font-size: 20px;margin-bottom: 8upx;">￥{{item.price}}</view>
		</view>
		<!-- <view class="demo-shop">{{item.description}}</view> -->
					
				</view>
			</template>
		</u-waterfall>
	</view>
</template>

<script>
	export default {
		name: "water",
		props: {
			list: {
				type: Array,
				default () {
					return [];
				}
			},
		},
		data() {
			return {}
		},
		
		methods: {
			toDetail(id) {
				console.log(id)
				// this.$common.navigateTo('/pages/index/detail')
				uni.navigateTo({
					url:'/pages/index/detail?id='+id
				})
			},
		}
	}
</script>

<style lang="scss">
	.demo-warter {
		border-radius: 8px;
		margin: 5px;
		background-color: #ffffff;
		padding: 8px;
		position: relative;
	}

	.demo-image {
		width: 100%;
		border-radius: 4px;
	}

	.demo-title {
		font-size: 30rpx;
		margin-top: 5px;
		color: $u-main-color;
		word-break: break-all;
		font-weight: 700;
	}

	.demo-tag {
		display: flex;
		flex-direction: row;
		align-items: center;
		flex-wrap: wrap;
		margin-top: 5px;
	}

	.demo-tag-owner {
		background-color: $u-type-error;
		color: #ffffff;
		display: flex;
		align-items: center;
		padding: 4rpx 14rpx;
		border-radius: 50rpx;
		font-size: 20rpx;
		line-height: 1;
	}

	.demo-tag-text {
		border: 1px solid $u-type-success;
		color: $u-type-success;
		margin-right: 10rpx;
		margin-bottom: 10rpx;
		border-radius: 50rpx;
		line-height: 1;
		padding: 8rpx 14rpx;
		display: flex;
		align-items: center;
		border-radius: 50rpx;
		font-size: 20rpx;
	}

	.demo-price {
		font-size: 30rpx;
		color: $u-type-error;
		margin-top: 5px;
	}

	.demo-shop {
		font-size: 22rpx;
		color: $u-tips-color;
	}
	
	.u-close {
		position: absolute;
		top: 32rpx;
		right: 32rpx;
	}
</style>
