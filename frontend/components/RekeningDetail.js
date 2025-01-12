import axios from 'axios'
import { useState, useEffect } from 'react'

const RekeningDetail = ({ id }) => {
    const [rekening, setRekening] = useState(null)
    const [transaksis, setTransaksis] = useState([])

    useEffect(() => {
        axios.get(`/api/rekening/${id}`)
            .then(response => {
                setRekening(response.data)
            })

        axios.get(`/api/transaksi?rekeningId=${id}`)
            .then(response => {
                setTransaksis(response.data)
            })
    }, [id])

    if (!rekening) return <div>Loading...</div>

    return (
        <div>
            <h2>{rekening.namaPemilik}</h2>
            <p>Nomor Rekening: {rekening.nomorRekening}</p>
            <p>Saldo: {rekening.saldo}</p>
            <p>Tanggal Pembuatan: {new Date(rekening.tanggalPembuatan).toLocaleString()}</p>

            <h3>Transaksi</h3>
            <ul>
                {transaksis.map(transaksi => (
                    <li key={transaksi.id}>
                        {transaksi.jenisTransaksi} - {transaksi.jumlahTransaksi} - {new Date(transaksi.tanggalTransaksi).toLocaleString()}
                    </li>
                ))}
            </ul>
        </div>
    )
}

export default RekeningDetail
