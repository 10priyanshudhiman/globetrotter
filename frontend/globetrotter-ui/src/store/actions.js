
export function updateUser(user) {
    return {
      type: 'UPDATE_USER',
      user,
    };
}

export function userActionResult(userActionResult) {
    return {
        type: 'UPDATE_USER_ACTION_RESPONSE',
        userActionResult
    }
}