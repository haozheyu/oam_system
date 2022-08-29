<template>
	<view>
		<u-navbar :background="background" back-text="返回" :back-text-style="{'color':'#fff'}" backIconColor="#fff">
			<view class="search-wrap">
				<!-- 如果使用u-search组件，必须要给v-model绑定一个变量 -->
				<u-search :show-action="showAction" inputAlign="center" :searchIconStyle="searchIconStyle" height="56" :action-style="{color: '#fff'}"></u-search>
			</view>
		</u-navbar>
		
		<water :list="flowList"></water>
		
		<u-gap height="30" bgColor="#F7F7F7"></u-gap>
		<!-- top s -->
		<u-back-top :scrollTop="scrollTop" mode="circle" bottom="200" right="40" top="600" :icon="icon"
			:icon-style="iconStyle" :tips="tips" :custom-style="customStyle">
		</u-back-top>
	</view>
</template>

<script>
	import water from '@/components/water';
	export default {
		components:{water},
		data() {
			return {
				background: {
					'background-image': 'linear-gradient(45deg, rgb(28, 187, 180), rgb(141, 198, 63))'
				},
				showAction: false,
				searchIconStyle: {
					'display':'none'
				},
				list:[
					{image:'/static/pic1.jpg'},
					{image:'/static/pic2.jpg'},
					{image:'/static/pic3.jpg'},
					{image:'/static/pic4.jpg'},
				],
				// top
				flowList: [],
				page: 1,
				is_loading: true,
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
		onLoad() {
			this.addRandomData()
		},
		methods:{
			
			addRandomData() {
				for (let i = 0; i < 10; i++) {
					let index = this.$u.random(0, this.list.length - 1);
					// 先转成字符串再转成对象，避免数组对象引用导致数据混乱
					let item = JSON.parse(JSON.stringify(this.list[index]));
					item.id = this.$u.guid();
					this.flowList.push(item);
				}
			},
			onReachBottom() {
				if (this.is_loading) {
					this.page++;
					this.addRandomData();
				}
			},
			// top
			onPageScroll(e) {
				this.scrollTop = e.scrollTop;
			},
		}
	}
</script>

<style lang="scss" scoped>
	.search-wrap {
		margin: 0 30rpx 0 10rpx;
		flex: 1;
	}
	.demo-warter {
		border-radius: 8px;
		margin: 5px;
		background-color: #ffffff;
		padding: 8px;
		position: relative;
	}
	
	.u-close {
		position: absolute;
		top: 32rpx;
		right: 32rpx;
	}
	
	.demo-img-wrap {}
	
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
		padding: 4rpx 14rpx;
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
</style>
