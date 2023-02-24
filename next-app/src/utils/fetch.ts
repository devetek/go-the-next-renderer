// import "server-only";
import axios from 'axios';

const fetchData = async (mockName: string = "", url?: string) => {
    if (mockName) {
        console.log("Fetch to mock:", mockName)
        // use dynamic import json https://javascript.info/modules-dynamic-imports
        const object = await import(`./../mockup/${mockName}.json`);
        return object.default;
    }

    const targetURL = url ? url : 'https://mocki.io/v1/4459ceb7-791b-4948-87e1-db598b805587';

    // real fetch when real access
    try {
        const response = await axios.get(targetURL);
        return response?.data || null;
    } catch (error) {
        console.error(error);
        return null;
    }
}

export default fetchData;