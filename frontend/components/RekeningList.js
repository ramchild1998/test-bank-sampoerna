import api from '../services/api'
import { useState, useEffect } from 'react'
import Link from 'next/link'

const RekeningList = () => {
    const [rekenings, setRekenings] = useState([])

    useEffect(() => {
        api.get('/rekening')
            .then(response => {
                setRekenings(response.data)
            })
    }, [])

    return (
        <div>
            <h2>Daftar Rekening</h2>
            <Link href="/rekening/new">
                <a>Tambah Rekening Baru</a>
            </Link>
            <ul>
                {rekenings.map(rekening => (
                    <li key={rekening.id}>
                        <Link href={`/rekening/${rekening.id}`}>
                            <a>{rekening.namaPemilik} - {rekening.nomorRekening}</a>
                        </Link>
                    </li>
                ))}
            </ul>
        </div>
    )
}

export default RekeningList
