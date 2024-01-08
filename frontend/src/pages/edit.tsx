// pages/edit.tsx
import { useState, useEffect } from 'react';
import axios from 'axios';

interface KeyValuePairs {
    [key: string]: string;
}

export default function Edit() {
    const [keyValuePairs, setKeyValuePairs] = useState<KeyValuePairs>({});

    useEffect(() => {
        // Fetch key-value pairs from the server
        const fetchData = async () => {
            try {
                const response = await axios.get<KeyValuePairs>('http://localhost:8080/key-value-pairs');
                setKeyValuePairs(response.data);
            } catch (error) {
                console.error(error);
            }
        };

        fetchData();
    }, []); // Run only once on component mount

    // Implement your logic for rendering and editing key-value pairs here

    return (
        <div>
            <h1>Edit Key-Value Pairs</h1>
            {/* Display and edit key-value pairs using keyValuePairs state */}
            {Object.entries(keyValuePairs).map(([key, value]) => (
                <div key={key}>
                    <input
                        type="text"
                        value={key}
                        // Handle changes to the key input
                    />
                    <input
                        type="text"
                        value={value}
                        // Handle changes to the value input
                    />
                </div>
            ))}
        </div>
    );
}
