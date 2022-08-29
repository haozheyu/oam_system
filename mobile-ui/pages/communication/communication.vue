<template>
	<view class="home">
		<u-navbar title="广场" :is-back="false" :background="background" titleColor="#fff"></u-navbar>
		<view class="card-bottom" style="margin-top: -92upx;">
			<!-- 顶部分页栏 -->
			<!-- <view class="top-tab">
				<view :class="['tab-item flex-center', activeTab == index ? 'active' : '']" @tap="handleTab(index)" v-for="(item, index) in tabList" :key="index">{{ item.title }}</view>
			</view> -->
				<view class="scroll-wrapper">
					<view v-if="activeTab == 0">
						<view class="margin-bottom" v-for="(item, index) in cardList" :key="index" >
							<y-DiaryItem :obj="item" />
						</view>
					</view>
					<view v-else>
						<view class="margin-bottom" v-for="(item, index) in rightList" :key="item.id">
							<y-DiaryItem :obj="item" />
						</view>
					</view>
					<y-LoadMore :status="loadMoreStatus" />
				</view>
			<!-- 右下角按钮 -->
			<y-Fab :bottom="140" :right="40" :btnList="fabList" @click="handleFab" />
		</view>
	</view>
</template>

<script>
var that;
import {food_list,note_list} from '@/api/food.js'

export default {
	data() {
		return {
			background: {
				'background-image': 'linear-gradient(45deg, rgb(28, 187, 180), rgb(141, 198, 63))'
			},
			startNum: 0,
			activeTab: 0,
			// tab的名称
			tabList: [
				{
					title: ''
				}
				
			],
			cardList: [],
			rightList: [],
			loadMoreStatus: 1, //0加载前，1加载中，2没有更多了
			//fab的设置
			fabList: [
				{
					bgColor: '#16C2C2',
					text: '帖子',
					fontSize: 28,
					color: '#fff'
				}
			]
		};
	},
	onShow() {
		
		
		
		// this.getNotesList()
	},
	onLoad() {
		// this.getFoodList()
		// this.getFoodList()
		// this.getNotesList()
		this.cardList=[]
		food_list({}).then(res=>{
			console.log(res)
			this.cardList=res.data
		})
	},
	onReachBottom() {

	},
	methods: {
		getFoodList(){
			food_list({}).then(res=>{
				console.log(res)
				this.cardList=res.data
			})
		},
		getNotesList(){
			note_list({}).then(res=>{
				console.log(res)
				this.rightList=res.data
			})
		},
		
		toDetails(id){
			uni.navigateTo({
				url: '../diary/details/details?id=' + id 
			})
		},
		toOther(id){
			uni.navigateTo({
				url: '../mine/other/other?id=' + id
			})
		},
		loadData(type) {
			
		},
		//下拉刷新
		onPulldownReresh() {
			// that.loadData('refresh');
			this.getFoodList()
		},
		handleTab(index) {
			this.activeTab = index;
			if(index===0){
				food_list({}).then(res=>{
					console.log(res)
					this.cardList=res.data
				})
			}else{
				note_list({}).then(res=>{
					console.log(res)
					this.rightList=res.data
				})
			}
		},
		//点击右下角tab按钮
		handleFab(e) {
			let index = e.index;
			if(index==0){
				uni.navigateTo({
					url: '../publish/publish'
				});
				
			}else{
				uni.navigateTo({
					url: '../publish-note/publish-note'
				});
			}
		
		}
	}
};
</script>

<style lang="less" scoped>
.home {
	padding-top: 120rpx;

	.top-barrage {
		width: 100%;
		height: 320rpx;
		overflow: hidden;
	}

	.card-bottom {
		width: 100%;

		.top-tab {
			display: flex;
			height: 120rpx;
			width: 100%;
			z-index: 100;
			background-color: #ffffff;

			.tab-item {
				flex: 1;
				color: #999;
				border-bottom: 4rpx solid #ececec;
			}

			.active {
				color: var(--mainColor);
				border-bottom: 4rpx solid var(--mainColor);
			}
		}
	}
}
</style>
