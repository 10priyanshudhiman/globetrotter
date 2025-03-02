const SERVER_API_URL=`http://${process.env.REACT_APP_BASE_URL}`

const URLS = {
    login: `${SERVER_API_URL}/user/details`,
    updateUser: `${SERVER_API_URL}/user/update`,

    gameDetails: `${SERVER_API_URL}/game/details`,
    verfiyAction: `${SERVER_API_URL}/game/verify`,
    getAllCountries: `${SERVER_API_URL}/game/country`

}

export default URLS