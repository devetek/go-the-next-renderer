import "server-only";
import axios from 'axios';
import userMock from './../mockup/user.json'

const fetchUser = async (isMock: boolean = false) => {
    if (isMock) {
        return userMock;
    }

    // real fetch when real access
    try {
        const response = await axios.get('https://mocki.io/v1/4459ceb7-791b-4948-87e1-db598b805587');
        return response?.data || null;
    } catch (error) {
        console.error(error);
        return null;
    }
}

export default fetchUser;