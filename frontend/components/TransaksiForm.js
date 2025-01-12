import api from '../services/api'
import { useFormik } from 'formik'

const TransaksiForm = ({ transaksi, rekeningId, onSave }) => {
    const formik = useFormik({
        initialValues: {
            nomorRekening: rekeningId,
            jenisTransaksi: transaksi ? transaksi.jenisTransaksi : 'debit',
            jumlahTransaksi: transaksi ? transaksi.jumlahTransaksi : 0
        },
        onSubmit: values => {
            if (transaksi) {
                api.put(`/transaksi/${transaksi.id}`, values)
                    .then(response => {
                        onSave(response.data)
                    })
            } else {
                api.post('/transaksi', values)
                    .then(response => {
                        onSave(response.data)
                    })
            }
        }
    })

    return (
        <form onSubmit={formik.handleSubmit}>
            <div>
                <label>Nomor Rekening</label>
                <input 
                    type="text" 
                    name="nomorRekening" 
                    value={formik.values.nomorRekening} 
                    onChange={formik.handleChange} 
                    disabled
                />
            </div>
            <div>
                <label>Jenis Transaksi</label>
                <select 
                    name="jenisTransaksi" 
                    value={formik.values.jenisTransaksi} 
                    onChange={formik.handleChange}
                >
                    <option value="debit">Debit</option>
                    <option value="kredit">Kredit</option>
                </select>
            </div>
            <div>
                <label>Jumlah Transaksi</label>
                <input 
                    type="number" 
                    name="jumlahTransaksi" 
                    value={formik.values.jumlahTransaksi} 
                    onChange={formik.handleChange} 
                />
            </div>
            <button type="submit">Save</button>
        </form>
    )
}

export default TransaksiForm
