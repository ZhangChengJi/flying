module.exports = {
	Login: {
		login: '登录',
		username: '用户名',
		password: '密码',
		loginSuccessfully: '登录成功',
		logoutSuccessfully: '登出成功',
		pleaseInputUsername: '请输入用户名',
		pleaseInputPassword: '请输入密码',
		invalidUsernameOrPassword: '用户名或密码错误',
		invalidAccess: '未登录或非法访问',
		tokenExpired: '授权token已过期',
		loginStatusFailed: '设置登录状态失败',
		getTokenFailed: '获取token失败',
		jwtInvalidationFailed: 'jwt作废失败',
		updatePassword: '修改密码',
		currentPassword: '当前密码',
		newPassword: '新密码',
		confirmPassword: '确认密码',
		pleaseInputCurrentPassword: '请输入当前密码',
		pleaseInputNewPassword: '请输入新密码',
		pleaseInputConfirmPassword: '请再次输入新密码',
		twoInputPassword: '两个输入不匹配!',
		lessSixNumber: '密码小于6位数',
		currentPasswordFailed: '当前密码错误',
		updatePasswordFailed: '更新密码错误',
		newCurrentPasswordNotSame: '新密码不能跟当前密码一致',
		logoutTitle: '是否退出系统, 是否继续?'



	},
	Header: {
		operation: '操作',
		envManagement: '环境管理',
		appManagement: '应用管理',
		languageSwitchButton: 'En',
		logout: '登出',
		changePassword: '修改密码',
	},
	Button:{
		cancel: '取消',
		create: '添加',
		update: '更新',
		delete: '删除',
		sync:   '同步',
		release: '发布',
		determine: '确定',
	},
	MessageTitle:{
		querySuccessfully: '查询成功',
		queryFailed: '查询失败',
		createSuccessfully: '创建成功',
		createFailed: '创建失败',
		updateSuccessfully: '修改成功',
		updateFailed: '修改失败',
		deleteSuccessfully: '删除成功',
		deleteFailed: '删除失败',
		connectionSuccess: '连接成功',
		connectionFailed: '连接失败',
		releaseSuccess: '发布成功',
		releaseFailed: '发布失败',
		syncSuccess: '同步成功',
		syncFailed: '同步失败',
	},
	Node: {
		node: '环境',
		nodeStatus: '环境状态',
		nodeManagement: '环境管理',
		addNode: '创建环境',
		updateNode: '修改环境',
		update: '修改',
		delete: '删除',
		testCon:'测试连接',
		areDelete:'您确定要删除环境:',
		confirmDelete: '删除前请确认是否还有应用使用该环境',
		nodeList: '环境列表',
		List:{
			conStatus: '连接状态',
			nodeName: '环境名称',
			nodeAddress: '环境地址',
			nodeAddressNotUpdate: '环境地址(地址不能修改)',
			nodeDescription: '环境说明',
			operation: '操作'
		},
        Error:{
			conError:'环境无法连接 请尽快修复',
			nodeConnectionFailed: '环境连接失败',
			nodeNameDuplicate: '环境名称被其他环境占用',
			nodeAddressDuplicate: '环境地址重复',
			nodeHistoryLegacy: '节点存在以往app,请清理数据库再使用',
			nodeNotFound: '未找到该环境',
			nodeNotDeleted: '不可以删除有%s应用在使用此环境',


		},
		Verification:{
			pleaseInputNodeName: '请输入环境名称',
			pleaseInputNodeAddress: '请输入环境地址'
		},
		placeholder:{
			nodeName: '必须是英文 比如dev pro fat',
			nodeAddress: 'flying-config 地址:端口/域名'
		}

	},
	App:{
		appProject:'应用项目',
		appManagement: '应用管理',
		addApp: '创建应用',
		updateApp: '修改应用',
		update: '修改',
		delete: '删除',
		areDelete:'是否确定要删除应用:',
		confirmDelete: '删除前请确认是否还有客户端使用该应用配置',
		appInfo: '应用信息',
		List:{
			appName: '应用名称',
			appId: 'AppId',
			useEnvironment: '使用环境',
			createTime: '创建时间',
			operation: '操作',
		},
		Verification:{
			pleaseInputAppId: '请输入AppId',
			pleaseInputAppName: '请输入应用名称',
			selectNode: '请至少选择一种/多种环境'
		},
		placeholder:{
			appId: '应用项目AppId',
			appName: '应用项目名称',
			selectNode: ' 环境无法连接，无法选择',
			updateCheckNode: '修改环境时，去掉原有勾选的环境，会清除该环境跟当前应用所有的信息，谨慎操作'
		},
		Error:{
			appExistsNode: '%s 环境已经存在此项目',
			appIdDuplicate: 'AppId:%s重复',
			appNotFind: '未找到应用项目',

		}
	},
	Namespace:{
		namespace: '命名空间',
		namespaceNodeError: '环境连接错误,无法查看',
		namespaceNodeDescription: '环境无法连接，namespace配置也无法查看 23333....',
		namespaceTestCaption: '您可以通过下面方式进行测试该环境是否正常：',
		namespaceTestConn: '测试该环境的连接地址',
		nodeAddress: '环境地址',
		isWrong: '地址是否有错',
		refreshDescription: '操作 > 环境管理 > 进行',
		refresh: '刷新',
		testConn: '测试连接更新状态',
		checkRunning: '查看该环境部署状态/运行情况是否正常',
		addNamespace: '创建命名空间',
		updateNamespace: '修改命名空间',
		syncConfig: '同步配置',
		List: {
			namespaceName: '命名空间名称',
			node: '环境',
			format: '格式',
			configuration: '配置信息'
		},
		Verification:{
			pleaseInputNamespaceName: '请输入命名空间名称',
			selectNode: '请至少选择一种/多种环境',
			selectFormat: '请选择配置格式',
			pleaseInputConfig: '请输入配置信息',

		},
		Error:{
			namespaceRepeat: '当前应用项目在%s环境已经存在此名称的命名空间',
			namespaceNotFind: '没有找到命名空间',
			namespaceNodeNotFind: '未找到到该namespace的环境',
			releaseOther: '此命名空间已被其他管理员发布'

		},
		Warning:{
			namespaceNoChange: '配置没有变动'
		},
		foldConfig:'展开/合上配置',
		releaseTag: '已发布',
		noReleaseTag: '未发布',
		lastReleaseTime: '最后发布时间: ',
		areDelete:'是否确定要删除命名空间:',
		confirmDelete: '删除后客户端会无法找到该namespace的配置',
        releaseTitle: '此命名空间配置发生变化,是否进行发布?',
		releaseName: '发布名称',
	},
	Release:{
		releaseSyncDescription: '要进行同步的命名空间在要被同步中环境中如果已经存在，会直接进行覆盖配置！如果没有会直接复制一份',
		releaseSyncDescription2: '命名空间只会同步到目前应用已有的环境(排除当前环境)！同步完成后记得进行到对于环境进行命名空间发布',
		selectSyncNode: '选择要同步的环境',
		releaseNamespaceSyncNotFound: '要被同步的命名空间没有找到',
		releaseNamespaceNodeSyncHistory: '应用项目在%s环境存在相同命名空间,是否进行覆盖?',
		releaseAppNoFind: '%s环境未找到当前应用项目'

	}

}