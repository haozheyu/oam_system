<template>
	<view>
		<u-navbar back-text="搜索" :borderBottom="false" :background="background" :back-text-style="{'color':'#fff'}"
			backIconColor="#fff"></u-navbar>
		<view class="background">
			<view class="home">
				<u-search inputAlign="center" placeholder="请输入关键词搜索" :searchIconStyle="searchIconStyle" shape="round"
					v-model="keyword" maxlength="32" :actionStyle="actionStyle" @custom="custom" @search="search"
					@change="change"></u-search>
			</view>
		</view>
		<view >
			<view class="u-flex u-row-between home">
				<view class="title-blod">历史记录</view>
				<view>
					<u-icon @click="delhistory" name="trash-fill" size="32"></u-icon>
				</view>
			</view>
			<view class="home" v-for="(item,index) in searchContentList" :key="index">
				<u-tag :text="item" type="error" mode="plain" class="tag"
					 shape="circle"></u-tag>
			</view>
		</view>
		<!-- 热搜榜 -->
		<!-- <view class="u-p-t-30">
			<view class="u-flex u-row-between home">
				<view class="title-blod">热门搜索</view>
				<view>
					<u-icon name="eye-fill" size="32"></u-icon>
				</view>
			</view>
			<view class="home">
				<u-tag text="清蒸鲈鱼" @click="click(item)" type="error" mode="plain" class="tag"
					v-for="(item,index) in 3" :key="index" shape="circle"></u-tag>
			</view>
		</view> -->
		
		<!-- top s -->
		<u-back-top :scrollTop="scrollTop" mode="circle" bottom="200" right="40" top="600" :icon="icon"
			:icon-style="iconStyle" :tips="tips" :custom-style="customStyle">
		</u-back-top>
	</view>
</template>

<script>
	import {create_search_content,search_list} from '@/api/search_content.js'
	
	export default {
		data() {
			return {
				background: {
					'background-image': 'linear-gradient(45deg, rgb(28, 187, 180), rgb(141, 198, 63))'
				},
				keyword: '',
				searchIconStyle: {
					'display': 'none'
				},
				actionStyle: {
					'color': '#fff'
				},
				searchContentList:[],
				page: 1,
				is_loading: true,
				// top
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
			this.searchList()
		},
		methods: {
			
			searchList(){
				let params={}
				search_list({}).then(res=>{
					console.log(res)
					this.searchContentList=res.data
					
				})
			},
			
			onReachBottom() {
				if (this.is_loading) {
					this.page++;
				}
			},
			delhistory() {
				this.$common.successToShow('清除成功');			
			},
			// top
			onPageScroll(e) {
				this.scrollTop = e.scrollTop;
			},
			custom(value) {
				console.log(value)
			},

			search(value) {
				console.log(value)
			},
			
			change(value) {
				console.log(value)
			},
		}
	}
</script>

<style lang="scss" scoped>
	.background{
		background-image: linear-gradient(45deg, rgb(28, 187, 180), rgb(141, 198, 63));
	}
	.tag {
		margin-right: 20rpx;
		display: inline-block;
		margin-bottom: 20rpx;
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
