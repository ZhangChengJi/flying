const success = 'success';
const error = 'error';
const warning='warning';
var util = require('util');
export function text(data) {
    switch (data.code) {
        case 401:
            return {title: '', msg: window.i18n.t('Login.invalidUsernameOrPassword'), type: error}
        case 10010:
            return {
                title: window.i18n.t('MessageTitle.createFailed'),
                msg: util.format(window.i18n.t('App.Error.appIdDuplicate'),data.msg),
                type: error
            }

        case 10011:
            return {
                title: window.i18n.t('MessageTitle.updateFailed'),
                msg: window.i18n.t('App.Error.appNotFind'),
                type: error
            }
        case 10012:
            return {
                title: window.i18n.t('MessageTitle.createFailed'),
                msg: util.format(window.i18n.t('App.Error.appExistsNode'),data.msg),
                type: error
            }
        case 10020:
            return {
                title: window.i18n.t('MessageTitle.createFailed'),
                msg: window.i18n.t('Node.Error.nodeConnectionFailed'),
                type: error
            }
        case 10021:
            return {
                title: window.i18n.t('MessageTitle.createFailed'),
                msg: window.i18n.t('Node.Error.nodeNameDuplicate'),
                type: error
            }
        case 10022:
            return {
                title: window.i18n.t('MessageTitle.createFailed'),
                msg: window.i18n.t('Node.Error.nodeAddressDuplicate'),
                type: error
            }
        case 10023:
            return {
                title: window.i18n.t('MessageTitle.createFailed'),
                msg: window.i18n.t('Node.Error.nodeHistoryLegacy'),
                type: error
            }
        case 10024:
            return {
                title: window.i18n.t('MessageTitle.updateFailed'),
                msg: window.i18n.t('Node.Error.nodeAddressDuplicate'),
                type: error
            }
        case 10025:
            return {
                title: window.i18n.t('MessageTitle.updateFailed'),
                msg: window.i18n.t('Node.Error.nodeNameDuplicate'),
                type: error
            }
        case 10026:

            return {
                title: window.i18n.t('MessageTitle.deleteFailed'),
                msg: util.format(window.i18n.t('Node.Error.nodeNotDeleted'),data.msg),
                type: error
            }
        case 10030:
            return {
                title: window.i18n.t('MessageTitle.createFailed'),
                msg: util.format(window.i18n.t('Namespace.Error.namespaceRepeat'),data.msg),
                type: error
            }
        case 10031:
            return {
                title: window.i18n.t('MessageTitle.updateFailed'),
                msg: util.format(window.i18n.t('Namespace.Error.namespaceNotFind'),data.msg),
                type: error
            }
        case 10032:
            return {
                title: window.i18n.t('MessageTitle.deleteFailed'),
                msg: util.format(window.i18n.t('Namespace.Error.namespaceNotFind'),data.msg),
                type: error
            }
        case 10033:
            return {
                title: window.i18n.t('MessageTitle.releaseFailed'),
                msg: window.i18n.t('Namespace.Warning.namespaceNoChange'),
                type: error
            }
        case 10034:
            return {
                title: window.i18n.t('MessageTitle.updateFailed'),
                msg: window.i18n.t('Namespace.Warning.namespaceNoChange'),
                type: warning
            }
        case 10035:
            return {
                title: window.i18n.t('MessageTitle.updateFailed'),
                msg: util.format(window.i18n.t('Namespace.Error.namespaceNodeNotFind'),data.msg),
                type: error
            }
        case 10036:
            return {
                title: window.i18n.t('MessageTitle.deleteFailed'),
                msg: window.i18n.t('Namespace.Error.namespaceNodeNotFind'),
                type: error
            }
        case 10040:
            return {
                title: window.i18n.t('MessageTitle.releaseFailed'),
                msg: window.i18n.t('Namespace.Error.releaseOther'),
                type: error
            }
        case 10041:
            return {
                title: window.i18n.t('MessageTitle.syncFailed'),
                msg: window.i18n.t('Release.releaseNamespaceSyncNotFound'),
                type: error
            }
        case 10050:
            return {
                title: window.i18n.t('MessageTitle.updateFailed'),
                msg: window.i18n.t('Login.currentPasswordFailed'),
                type: error
            }
        case 10051:
            return {
                title: window.i18n.t('MessageTitle.updateFailed'),
                msg: window.i18n.t('Login.updatePasswordFailed'),
                type: error
            }
        case 10052:
            return {
                title: window.i18n.t('MessageTitle.updateFailed'),
                msg: window.i18n.t('Login.newCurrentPasswordNotSame'),
                type: error
            }
        case 20010:
            return {
                title: window.i18n.t('MessageTitle.createSuccessfully'),
                msg: '',
                type: success
            }
        case 20011:
            return {
                title: window.i18n.t('MessageTitle.updateSuccessfully'),
                msg: '',
                type: success
            }
        case 20012:
            return {
                title: window.i18n.t('MessageTitle.deleteSuccessfully'),
                msg: '',
                type: success
            }
        case 20013:
            return {
                title: window.i18n.t('MessageTitle.connectionSuccess'),
                msg: '',
                type: success
            }
        case 20014:
            return {
                title: window.i18n.t('MessageTitle.releaseSuccess'),
                msg: '',
                type: success
            }
        case 20015:
            return {
                title: window.i18n.t('MessageTitle.syncSuccess'),
                msg: '',
                type: success
            }
        case 40100:
            return {
                title: window.i18n.t('Login.invalidUsernameOrPassword'),
                msg: '',
                type: error
            }
        case 40101:
            return {
                title: window.i18n.t('Login.invalidAccess'),
                msg: '',
                type: error
            }
        case 40102:
            return {
                title: window.i18n.t('Login.tokenExpired'),
                msg: '',
                type: error
            }
        case 40103:
            return {
                title: window.i18n.t('Login.loginStatusFailed'),
                msg: '',
                type: error
            }
        case 40104:
            return {
                title: window.i18n.t('Login.getTokenFailed'),
                msg: '',
                type: error
            }
        case 40105:
            return {
                title: window.i18n.t('Login.jwtInvalidationFailed'),
                msg: '',
                type: error
            }
        case 50010:
            return {
                title: window.i18n.t('MessageTitle.createFailed'),
                msg: '',
                type: error
            }
        case 50011:
            return {
                title: window.i18n.t('MessageTitle.updateFailed'),
                msg: '',
                type: error
            }
        case 50012:
            return {
                title: window.i18n.t('MessageTitle.deleteFailed'),
                msg: '',
                type: error
            }
        case 50013:
            return {
                title: window.i18n.t('MessageTitle.connectionFailed'),
                msg: '',
                type: error
            }
        case 50014:
            return {
                title: window.i18n.t('MessageTitle.releaseFailed'),
                msg: '',
                type: error
            }
        case 50015:
            return {
                title: window.i18n.t('MessageTitle.syncFailed'),
                msg: '',
                type: error
            }
        default:
            return ''
    }
}
