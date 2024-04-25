import React, {useState} from 'react';
import Axios from "axios";
import {Table} from "antd";

function App() {

    const [students, setStudents] = useState([])

    React.useEffect(() => {
        const fetchStudents = async () => {
            try {
                const {data} = await Axios.get('http://localhost:9999/student');
                console.log(data.data)
                const totalScore = data.data.reduce((prev, item) => {
                    const {score} = item;

                    if (score > 90 && score <= 100) {
                        item.class = 'A';
                    } else if (score > 80 && score <= 90) {
                        item.class = 'B';
                    } else if (score > 70 && score <= 80) {
                        item.class = 'C';
                    } else if (score > 60 && score <= 70) {
                        item.class = 'D';
                    } else {
                        item.class = 'E';
                    }

                    prev.push(item);
                    return prev;

                }, []);

                setStudents(totalScore);

            } catch (error) {
                console.error('Failed to fetch students:', error);
            }
        };

        fetchStudents();
    }, []);

    const columns = [
        {
            title: 'ID',
            dataIndex: 'id',
            key: 'id'
        },
        {
            title: 'Name',
            dataIndex: 'name',
            key: 'name'
        },
        {
            title: 'Age',
            dataIndex: 'age',
            key: 'age'
        },
        {
            title: 'Score',
            dataIndex: 'score',
            key: 'score'
        },
        {
            title: 'Class',
            dataIndex: 'class',
            key: 'class'
        },
    ]

    return (
        <div>
            <h1>Hello React</h1>
            <Table columns={columns} dataSource={students} rowKey={"id"}></Table>
        </div>
    )
}

export default App;