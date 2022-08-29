<template>
	<view class="page">
	<u-navbar title="个人信息" :is-back="true" :background="background" titleColor="#fff"></u-navbar>
	
	<!-- 基本信息 -->
	<view class="publish-demand-wrapper">
		<view style="display: flex;flex-direction: row;align-items: center;
		margin-left: 30upx;border-bottom: 1upx solid #F8F8F8;height: 96upx;justify-content: space-between;">
			<view style="display: flex;align-items: center;">
				<image style="width: 16upx;height: 16upx;" src="" mode="">
				</image>
				<span style="color: #191919;font-size: 28upx;margin-left: 18upx;">用户头像</span>
			</view>
			
			<img :src="userInfoForm.avatar" alt="" style="height: 60upx;width: 60upx;border-radius: 30upx;margin-right: 60upx;" @click="uploadImgLocal">
			<!-- <u-icon name="arrow-right" color="#969799" size="28"></u-icon> -->
		</view>
		
		<view style="display: flex;flex-direction: row;align-items: center;
		margin-left: 30upx;border-bottom: 1upx solid #F8F8F8;height: 96upx;">
			<view style="display: flex;align-items: center;">
				<image style="width: 16upx;height: 16upx;" src="" mode="">
				</image>
				<span style="color: #191919;font-size: 28upx;margin-left: 18upx;">用户昵称</span>
			</view>
	
			<input type="text" style="text-align: right;width: 70%;color: #191919;" maxlength="15" v-model="userInfoForm.nickname"
				placeholder="请输入用户昵称" />
	
		</view>
		
		<view style="display: flex;flex-direction: row;align-items: center;
		margin-left: 30upx;border-bottom: 1upx solid #F8F8F8;height: 96upx;">
			<view style="display: flex;align-items: center;">
				<image style="width: 16upx;height: 16upx;" src="" mode="">
				</image>
				<span style="color: #191919;font-size: 28upx;margin-left: 18upx;">手机号码</span>
			</view>
			
			<input type="text" style="text-align: right;width: 70%;color: #191919;" maxlength="15" v-model="userInfoForm.mobile"
				placeholder="请输入手机号码" />
			
		</view>
		
		<view style="display: flex;flex-direction: row;align-items: center;
		margin-left: 30upx;border-bottom: 1upx solid #F8F8F8;height: 96upx;">
			<view style="display: flex;align-items: center;">
				<image style="width: 16upx;height: 16upx;" src="" mode="">
				</image>
				<span style="color: #191919;font-size: 28upx;margin-left: 18upx;">个人密码</span>
			</view>
			
			<input type="text" style="text-align: right;width: 70%;color: #191919;" maxlength="15" v-model="userInfoForm.password"
				placeholder="请输入登录密码" />
			
		</view>
	
		
		
		
		
		
	</view>
	
	
		
		


		<view style="position: fixed;bottom: 0upx;margin-bottom: 68upx;padding: 0 30upx;" @click="publishDemand">
			<view style="width: 690upx;height: 80upx;background: #6BC362;border-radius: 40upx;
			text-align: center;line-height: 80upx;">确认修改</view>
		</view>

	</view>
</template>

<script>
	
	import {food_type_list} from '@/api/food_type.js'
	import {food_category_list} from '@/api/food_category.js'
	import {publish_food} from '@/api/food.js'
	import {user_info,update_user} from '@/api/user.js'
	import {
		set
	} from '@/utils/local.js'

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
				avatar: "",
				description: "",
				title: "",
				cookTime: "",
				selectCategroy: "",
				isAuth: 2,
				isTechAuth: 2,
				categoryId: '',
				estimate: 0,
				techType: 0,
				estimateList: [],
				foodTypeList: [],
				categoryList: [],
				estimateAmountId: 1,
				name:'',
				techTypeId: 1,
				foodTypeId:'',
				selectFoodType:'',
				date: getDate({
					format: true
				}),
				startDate: getDate('start'),
				endDate: getDate('end'),
				uploadImgList:[],
				imgList:[],
				userInfoForm:{}
			}
		},

		onShow() {
			
			// this.getFoodTypeList();
			// this.getFoodCategoryList();
			

		},
		onLoad() {
			user_info({}).then(res=>{
				console.log(res)
				this.userInfoForm=res.data
			})
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
						this.userInfoForm.avatar=tempFilePaths[0]
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
				
				let params = this.userInfoForm;
				
				console.log(params)
				uni.showLoading({
					title:'修改中~',
					mask:true
				})
				set('userId',params.id)
				
				set('nickname',params.nickname)
				this.uploadImgs().then(() =>{
					params.avatar=this.imgList[0]
					set('avatar',params.avatar)
					update_user(params).then(res => {
						if (res.status === 200) {
							uni.hideLoading()
							uni.showToast({
								title:'修改成功'
							})
							uni.switchTab({
								url:'/pages/mine/index'
							})
						
							// uni.switchTab({
							// 	url:'../communication/communication'
							// })
						}
					})
				})
			},

			bindDateChange: function(e) {
				this.date = e.detail.value;
			},
			getFoodTypeList() {
				food_type_list().then(res => {
					if (res.status === 200) {
						console.log(res)
						this.foodTypeList = res.data
					} 

				})
			},
			getFoodCategoryList() {
				food_category_list().then(res => {
					if (res.status === 200) {
						console.log(res)
						this.categoryList = res.data
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
			const index = e.target.value
			this.selectFoodType = this.foodTypeList[index].name
			this.foodTypeId = this.foodTypeList[index].id
			},
			bindTechTypePickerChange: function(e) {
				const index = e.target.value
				this.selectCategroy = this.categoryList[index].name
				this.categoryId = this.categoryList[index].id
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