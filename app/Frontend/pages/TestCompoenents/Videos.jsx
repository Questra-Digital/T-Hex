import React, { useEffect, useState } from "react";
import { gql } from "@apollo/client";
import client from "../../appolo-client";
import VideoPlayer from "../InsidePages/VideoPlayer";

const Videos = () => {
    const [dataToPass, setDataToPass] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    const handleResultPageLoad = async () => {
        try {
            const { data } = await client.query({
                query: gql`
                    query Tests {
                        tests {
                            id
                            username
                            project
                            type
                            urlid
                            testfile
                            status
                            duration
                        }
                    }
                `,
            });

            const filteredData = data.tests.filter(test => test.type === 'video');
            setDataToPass(filteredData);
            setLoading(false);
        } catch (error) {
            console.error("Error fetching data:", error);
            setError(error);
            setLoading(false);
        }
    };

    useEffect(() => {
        handleResultPageLoad();
    }, []); 

    if (loading) {
        return <div>Loading...</div>;
    }

    if (error) {
        return <div>Error: {error.message}</div>;
    }

    return (
        <>
            <VideoPlayer data={dataToPass}/>
        </>
    );
};

export default Videos;
