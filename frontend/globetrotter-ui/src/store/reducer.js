

const ApiReducer = (
    state = {
        user: {},
        userActionResult: {}
    },
    action,
) => {
    switch (action.type) {
        case 'UPDATE_USER':
            return { ...state, user: action.user }
        case 'UPDATE_USER_ACTION_RESPONSE':
            return { ...state, userActionResult: action.userActionResult }

        default:
            return state;
    }
};

export default ApiReducer;