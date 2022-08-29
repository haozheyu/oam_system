<template>
	<view class="page">
	<u-navbar title="发布笔记" :is-back="true" :background="background" titleColor="#fff"></u-navbar>
		<view style="background: #FFFFFF;margin-top: 30upx;">
			<textarea style="padding: 30upx 0 0 30upx;color: #191919;height: 272upx;width: 95%;" maxlength="1000" v-model="description" value=""
				placeholder="描述下你的笔记吧~(必填)" />

					
			<view style="margin-left: 30upx;">
				<view style="color: #B7B8BB;font-size: 28upx;">图片描述(必填)</view>
			</view>
	
			<view style="display: flex;flex-direction: row;align-items: center;">
				<image style="width: 160upx;height: 160upx;margin-top: 26upx;margin-bottom: 30upx;margin-left: 30upx;"
					src="../../static/img/publish-img-icon.png" mode="" @click="uploadImgLocal"></image>
				<view style="display: flex;flex-direction: row;position: relative;" v-for="(item,index) in uploadImgList" :key="index">
					
					
					<image :src="item" @tap="previewImg(index)" style="width: 160upx;height: 160upx;border-radius: 12upx;margin-left: 16upx;" mode="aspectFill"></image>
					<view style="background: #000000;border-radius: 16upx;width: 32upx;height: 32upx;line-height: 32upx;
					position: absolute;left: 75%;top: 8%;
					text-align: center;opacity:0.6" @click="removeImg(index)">
						<span style="color: #FFFFFF;width: 14upx;height: 14upx;">×</span>
					</view>
				</view>	
				
			</view>

		</view>

		<!-- 基本信息 -->
		<view class="publish-demand-wrapper">
			<view style="display: flex;flex-direction: row;align-items: center;
			margin-left: 30upx;border-bottom: 1upx solid #F8F8F8;height: 96upx;">
				<view style="display: flex;align-items: center;">
					<image style="width: 16upx;height: 16upx;" src="" mode="">
					</image>
					<span style="color: #191919;font-size: 28upx;margin-left: 18upx;">笔记名称</span>
				</view>

				<input type="text" style="text-align: right;width: 70%;color: #191919;" maxlength="15" v-model="title"
					placeholder="给笔记起个名称吧~" />

			</view>

		</view>




		<view style="position: fixed;bottom: 0upx;margin-bottom: 68upx;padding: 0 30upx;" @click="publishDemand">
			<view style="width: 690upx;height: 80upx;background: #6BC362;border-radius: 40upx;
			text-align: center;line-height: 80upx;">发布</view>
		</view>

	</view>
</template>

<script>
	
	import {create_notes} from '../../api/food.js'

	function getDate(type) {
		const date = new Date();

		let year = date.getFullYear();
		let month = date.getMonth() + 1;
		let day = date.getDate();

		if (type === 'start') {
			year = year - 10;
		} else if (type === 'end') {
			year = year + 10;
		}
		month = month > 9 ? month : '0' + month;
		day = day > 9 ? day : '0' + day;

		return `${year}-${month}-${day}`;
	}
	export default {
		data() {
			return {
				background: {
					'background-image': 'linear-gradient(45deg, rgb(28, 187, 180), rgb(141, 198, 63))'
				},
				isLogin:false,
				mobile: "",
				email: "",
				wechatAccount: "",
				qqAccount: "",
				description: "",
				title: "",
				isAuth: 2,
				isTechAuth: 2,
				estimate: 0,
				techType: 0,
				estimateList: [],
				techTypeList: [],
				estimateAmountId: 1,
				techTypeId: 1,
				date: getDate({
					format: true
				}),
				startDate: getDate('start'),
				endDate: getDate('end'),
				uploadImgList:[],
				imgList:[]
			}
		},

		onShow() {
			

	

		},
		methods: {
			previewImg(index) {
				console.log("index:"+index)
				let _this = this;
				let imgsArray = [];
				for (let i = 0; i < this.uploadImgList.length; i++) {
						imgsArray.push(this.uploadImgList[i]);
					
				}
				uni.previewImage({
					current: index,
					urls: imgsArray,
					indicator: 'number',
					loop: true,
					longPressActions:{
						itemList:['保存图片'],
						success(res) {
							console.log(res)
							uni.downloadFile({
								url:this.uploadImgList[index],
								success(result) {
									uni.saveImageToPhotosAlbum({
										filePath:result.tempFilePath,
										success() {
											uni.showToast({
												title:'保存成功'
											})
										}
									})
								}
							})
						}
					}
				});
			},
			removeImg(index){
				console.log("index:"+index)
				var temArray=[]
				for (var i = 0; i < this.uploadImgList.length; i++) {
					if(i !==index){
						temArray.push(this.uploadImgList[i])
					}
				}
				this.uploadImgList=temArray
				
			},
			uploadImgLocal() {
				uni.chooseImage({
					success: (chooseImageRes) => {
						const tempFilePaths = chooseImageRes.tempFilePaths;
						console.log(tempFilePaths)
						this.uploadImgList=tempFilePaths
					}
				})
			},
			uploadImgs(){
				const arr = []
				
				const _this = this
				for (var i = 0; i < this.uploadImgList.length; i++) {
					const a = new Promise(reslut => {
						uni.uploadFile({
							url: 'https://apibyte.wkhelpme.com/mars/upload/v1/img',
							filePath:this.uploadImgList[i],
							name:'file',
							success(res) {
								_this.imgList.push(JSON.parse(res.data).data)
								reslut()
							}
						})
					})
					arr.push(a)
				}
				return Promise.all(arr)
			},
			publishDemand() {
				
				let params = {}
				
				if (this.description === null || this.description === '') {
					uni.showToast({
						title: '请输入笔记描述',
						icon:'none',
						duration:1000
					})
					return
				}
				if (this.title === null || this.title === '') {
					uni.showToast({
						title: '请输入笔记标题',
						icon:'none',
						duration:1000
					})
					return
				}

				params.description = this.description
				params.title = this.title
				
				uni.showLoading({
					title:'发布中~',
					mask:true
				})
				this.uploadImgs().then(() =>{
					params.imgUrl=this.imgList
					create_notes(params).then(res => {
						if (res.status === 200) {
							uni.hideLoading()
							uni.showToast({
								title:"笔记发布成功",
								duration:2000
							})
							uni.switchTab({
								url:'../communication/communication'
							})
						}
					})
				})
			},

			bindDateChange: function(e) {
				this.date = e.detail.value;
			},
			getEstimateList() {
				estimate_list().then(res => {
					if (res.status === 200) {
						this.estimateList = res.data
					} else if (res.status === 401) {
						uni.navigateTo({
							url: '../login/login'
						})
					}

				})
			},
			getTechTypeList() {
				techType_pull().then(res => {
					if (res.status === 200) {
						this.techTypeList = res.data
					} else if (res.status === 401) {
						uni.navigateTo({
							url: '../login/login'
						})
					}

				})
			},

			linkPublishSuccess() {
				uni.navigateTo({
					url: "../publish-demand-success/publish-demand-success"
				})
			},
			bindPickerChange: function(e, estimateList) {
				this.estimate = e.target.value
				this.estimateAmountId = (estimateList[this.estimate].id)
				console.log(this.estimateAmountId)
			},
			bindTechTypePickerChange: function(e, techTypeList) {
				this.techType = e.target.value
				this.techTypeId = (techTypeList[this.techType].id)
				console.log(this.techTypeId)
			}
			
		}
	}
</script>

<style>
	.page {
		color: #F8F8F8;
	}

	.publish-demand-wrapper {
		display: flex;
		flex-direction: column;
		margin-top: 24upx;
		background: #FFFFFF;
	}
</style>