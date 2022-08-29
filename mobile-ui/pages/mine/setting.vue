<template>
	<view>
		<u-navbar back-text="设置" :borderBottom="false" :background="background" :back-text-style="{'color':'#fff'}"
			backIconColor="#fff"></u-navbar>
		<view class="">
			<u-cell-group>
				<u-cell-item title="头像" valueImage="/static/pic.jpg"></u-cell-item>
				<u-cell-item title="昵称" value="昵称"></u-cell-item>
				<u-cell-item title="性别" value="男"></u-cell-item>
				<u-cell-item title="国家" value="中国"></u-cell-item>
				<u-cell-item title="语言" value="中文"></u-cell-item>
			</u-cell-group>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				background: {
					'background-image': 'linear-gradient(45deg, rgb(28, 187, 180), rgb(141, 198, 63))'
				},
				userinfo:{}
			}
		},
		onLoad() {
			
		},
		methods: {

		}
	}
</script>

<style lang="scss" scoped>
</style>
