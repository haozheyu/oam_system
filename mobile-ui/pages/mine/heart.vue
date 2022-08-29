<template>
	<view>
		<u-navbar back-text="收藏" :borderBottom="false" :background="background" :back-text-style="{'color':'#fff'}"
			backIconColor="#fff"></u-navbar>
		<water :list="flowList"></water>
	</view>
</template>

<script>
	import water from '@/components/water'
	export default {
		components:{water},
		data() {
			return {
				background: {
					'background-image': 'linear-gradient(45deg, rgb(28, 187, 180), rgb(141, 198, 63))'
				},
				list:[
					{image:'/static/pic1.jpg'},
					{image:'/static/pic2.jpg'},
					{image:'/static/pic3.jpg'},
					{image:'/static/pic4.jpg'},
				],
				flowList:[]
			}
		},
		onLoad() {
			this.addRandomData()
		},
		methods: {
			addRandomData() {
				for (let i = 0; i < 10; i++) {
					let index = this.$u.random(0, this.list.length - 1);
					// 先转成字符串再转成对象，避免数组对象引用导致数据混乱
					let item = JSON.parse(JSON.stringify(this.list[index]));
					item.id = this.$u.guid();
					this.flowList.push(item);
				}
			},
		}
	}
</script>

<style lang="scss">
	page{
		background-color: #F7F7F7;
	}
</style>
