module.exports = {
    Login: {
        login: 'Login',
        username: 'Username',
        password: 'Password',
        loginSuccessfully: 'Login successfully',
        logoutSuccessfully: 'Logout successfully',
        pleaseInputUsername: 'Please input username',
        pleaseInputPassword: 'Please input password',
        invalidUsernameOrPassword: 'invalid username or password',
        invalidAccess: 'No login or illegal access',
        tokenExpired: 'The authorization token has expired',
        loginStatusFailed: 'Failed to set login status',
        getTokenFailed: 'get token failed',
        jwtInvalidationFailed: 'jwt invalidation failed',
        updatePassword: 'Update password',
        currentPassword: 'Current password',
        newPassword: 'New password',
        confirmPassword: 'Confirm password',
        pleaseInputCurrentPassword: 'Please input current password',
        pleaseInputNewPassword: 'Please input new password',
        pleaseInputConfirmPassword: 'Please input the new password again',
        twoInputPassword: 'Two inputs don t match!',
        lessSixNumber: 'Password is less than 6 digits',
        currentPasswordFailed: 'The current password is wrong',
        updatePasswordFailed: 'Update password Failed',
        newCurrentPasswordNotSame: 'The new password cannot be the same as the current password',
        logoutTitle: 'Exit the system, continue?'
    },
    Header: {
        operation: 'Operation',
        envManagement: 'Environmental management',
        appManagement: 'Application management',
        languageSwitchButton: '中',
        logout: 'logout',
        changePassword: 'modify password',
    },
    Button:{
        cancel: 'Cancel',
        create: 'Create',
        update: 'Update',
        delete: 'Delete',
        sync:   'Sync',
        release: 'Release',
        determine: 'OK',
    },
    MessageTitle:{
        querySuccessfully: 'Search successfully',
        queryFailed: 'Search failed',
        createSuccessfully: 'Create successfully',
        createFailed: 'Create failed',
        updateSuccessfully: 'Update successfully',
        updateFailed: 'Update failed',
        deleteSuccessfully: 'Deleted successfully',
        deleteFailed: 'Delete failed',
        connectionSuccess: 'Connection succeeded',
        connectionFailed: 'Connection failed',
        releaseSuccess: 'Release succeeded',
        releaseFailed: 'Release failed',
        syncSuccess: 'Synchronize succeeded',
        syncFailed: 'Synchronize failed'
    },
    Status:{
        createFailed: 'Create Failed',
        createSuccessfully: 'Create Successfully'
    },
    Node: {
        node: 'Environmental',
        nodeStatus: 'Environmental status',
        nodeManagement: 'Environmental management',
        addNode: 'Add environmental',
        updateNode: 'update environmental',
        update: 'update',
        delete: 'delete',
        testCon:'Test connection',
        areDelete:'Are you sure you want to delete the environment:',
        confirmDelete: 'Please confirm whether there are any applications using the environment before deleting',
        nodeList: 'Environmental list',
        List:{
            conStatus: 'Connection status',
            nodeName: 'Environmental name',
            nodeAddress: 'Environmental Address',
            nodeAddressNotUpdate: 'Environment address (address cannot be modified)',
            nodeDescription: 'Environmental description',
            operation: 'operation'
        },
        Error:{
            conError:'environment cannot be connected. Please repair it as soon as possible',
            nodeConnectionFailed: 'environment failed connected',
            nodeNameDuplicate: 'Environment name is occupied by other environments',
            nodeAddressDuplicate: 'Duplicate environment address name',
            nodeHistoryLegacy: 'The node has past apps, please clean up the database and use it again',
            nodeUpdateNotFound: 'The environment was not found',
            nodeNotDeleted: 'It is not possible to delete the %s application that is using this environment'
        },
        Verification:{
            pleaseInputNodeName: 'Please input environment name',
            pleaseInputNodeAddress: 'Please input environment address',
            nodeHistoryLegacy: 'The node has past apps, please clean up the database and use it again'

        },
        placeholder:{
            nodeName: 'Must be in English, such as dev pro fat',
            nodeAddress: 'flying-config address:port/domain name'
        }

    },
    App:{
        appProject: 'Application project',
        appManagement: 'Application management',
        addApp: 'Create Application',
        updateApp: 'Update Application',
        update: 'update',
        delete: 'delete',
        areDelete:'Are you sure you want to delete the application:',
        confirmDelete: 'Please confirm whether there are still clients using the application configuration before deleting',
        appInfo: 'Application information',
        List:{
            appName: 'Application Name',
            appId: 'AppId',
            useEnvironment: 'Use environment',
            createTime: 'Create time',
            operation: 'operation'
        },
        Verification:{
            pleaseInputAppId: 'Please input AppId',
            pleaseInputAppName: 'Please input Application Name',
            selectNode: 'Please select at least one/multiple environments'
        },
        placeholder:{
            appId: 'Application Project AppId',
            appName: 'Application project name',
            selectNode: 'The environment cannot be connected and cannot be selected',
            updateCheckNode: 'When modifying the environment, remove the original checked environment, it will clear all the information of the environment and the current application, operate with caution'
        },
        Error:{
            appExistsNode: 'This project already exists in %s environment',
            appIdDuplicate: 'AppId: Duplicate %s',
            appNotFind: 'No application found',

        }
    },
    Namespace:{
        namespace: 'namespace',
        namespaceNodeError: 'Environment connection error, unable to view',
        namespaceNodeDescription: 'The environment cannot be connected, and the namespace configuration cannot be viewed 23333....',
        namespaceTestCaption: 'You can test whether the environment is normal by the following methods：',
        namespaceTestConn: 'Test the connection address of the environment',
        nodeAddress: 'Environment address',
        isWrong: 'Is the address wrong?',
        refreshDescription: 'Operation> Environmental Management> Proceed',
        refresh: 'refresh',
        testConn: 'Test connection update status',
        checkRunning: 'Check whether the deployment status/operation of the environment is normal',
        addNamespace: 'Create namespace',
        updateNamespace: 'Update namespace',
        syncConfig: 'Synchronize Configuration',

        List: {
            namespaceName: 'Namespace name',
            node: 'Environmental',
            format: 'Format',
            configuration: 'Configuration'
        },
        Verification:{
            pleaseInputNamespaceName: 'Please input namespace name',
            selectNode: 'Please select at least one/multiple environments',
            selectFormat: 'Please select configuration format',
            pleaseInputConfig: 'Please input configuration',

        },
        Error:{
            namespaceRepeat: 'The current application project already has a namespace with this name in the %s environment',
            namespaceNotFind: 'No namespace found',
            namespaceNodeNotFind: 'The environment of the namespace was not found',
            releaseOther: 'This namespace has been published by another administrator'
        },
        Warning:{
            namespaceNoChange: 'The configuration has not changed'
        },
        foldConfig:'Open/fold configuration',
        releaseTag: 'Release',
        noReleaseTag: 'No release',
        lastReleaseTime: 'Last release time: ',
        areDelete:'Are you sure you want to delete the namespace:',
        confirmDelete: 'After deletion, the client will not be able to find the configuration of the namespace',
        releaseTitle: 'This namespace configuration has changed, do you want to publish it?',
        releaseName: 'Release name'
    },
    Release:{
        releaseSyncDescription: 'If the namespace to be synchronized already exists in the environment to be synchronized, it will be overwritten directly! If not, a copy will be made directly',
        releaseSyncDescription2: 'The namespace will only be synchronized to the existing environment of the current application (excluding the current environment)! After the synchronization is complete, remember to proceed to the namespace publishing for the environment',
        selectSyncNode: 'Please select sync environment',
        releaseNamespaceSyncNotFound: 'The namespace to be synchronized was not found',
        releaseNamespaceNodeSyncHistory: 'The application project has the same namespace in the %s environment, do you want to overwrite it?'

    }
}