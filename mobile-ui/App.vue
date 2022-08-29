<script>
	export default {
		onLaunch: function() {
			console.log('App Launch')
		},
		onShow: function() {
			console.log('App Show')
		},
		onHide: function() {
			console.log('App Hide')
		}
	}
</script>

<style lang="scss">
	@import "uview-ui/index.scss";
	@import 'colorui/main.css';
	@import 'colorui/icon.css';
	@import 'colorui/animation.css';
	
	/*每个页面公共css */
	@font-face {
		font-family: 'AliR';
		src: url(https://vkceyugu.cdn.bspapp.com/VKCEYUGU-matchbox/a003f470-a6fa-11ea-a30b-e311646dfaf2.otf);
	}
	
	@font-face {
		font-family: 'AliM';
		src: url(https://vkceyugu.cdn.bspapp.com/VKCEYUGU-matchbox/9ff65fe0-a6fa-11ea-a30b-e311646dfaf2.otf);
	}
	
	page {
		--mainColor: #435257;
		--activeColor: #36b39b;
	}
	
	body {
		font-family: 'AliR';
		background-color: #f5f8f9;
	}
	
	view {
		box-sizing: border-box;
	}
	
	.flex-center {
		display: flex;
		align-items: center;
		justify-content: center;
	}
	
	.mainColor {
		color: var(--mainColor);
	}
	
	.aColor {
		color: var(--activeColor);
	}
	
	.color-nine{
		color: #999999;
	}
	
	.main-btn {
		border-radius: 40upx;
		display: flex;
		align-items: center;
		justify-content: center;
		color: var(--mainColor);
		border: 1upx solid var(--mainColor);
		padding: 10rpx 40rpx;
	}
	
	.active-btn {
		color: #FFFFFF !important;
		background-color: var(--activeColor) !important;
		border: 1upx solid var(--activeColor) !important;
	}
	
	/* 点赞和评论 */
	.bottom-btn {
		display: flex;
	
		.btn-item {
			flex: 1;
			color: var(--mainColor);
	
			.img {
				width: 50rpx;
				height: 50rpx;
				margin-right: 20rpx;
			}
		}
	}
	
	
/************************************************************
** 请将全局样式拷贝到项目的全局 CSS 文件或者当前页面的顶部 **
** 否则页面将无法正常显示                                  **
************************************************************/

html {
  font-size: 16px;
}

body {
  margin: 0;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', 'Fira Sans',
    'Droid Sans', 'Helvetica Neue', 'Microsoft Yahei', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

view,
image,
text {
  box-sizing: border-box;
  flex-shrink: 0;
}

#app {
  width: 100vw;
  height: 100vh;
}

.flex-row {
  display: flex;
  flex-direction: row;
}

.flex-col {
  display: flex;
  flex-direction: column;
}

.justify-start {
  display: flex;
  justify-content: flex-start;
}

.justify-center {
  display: flex;
  justify-content: center;
}

.justify-end {
  display: flex;
  justify-content: flex-end;
}

.justify-evenly {
  display: flex;
  justify-content: space-evenly;
}

.justify-around {
  display: flex;
  justify-content: space-around;
}

.justify-between {
  display: flex;
  justify-content: space-between;
}

.items-start {
  display: flex;
  align-items: flex-start;
}

.items-center {
  display: flex;
  align-items: center;
}

.items-end {
  display: flex;
  align-items: flex-end;
}
</style>
