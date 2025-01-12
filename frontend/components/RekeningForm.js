import api from '../services/api'
import { useFormik } from 'formik'

const RekeningForm = ({ rekening, onSave }) => {
    const formik = useFormik({
        initialValues: {
            namaPemilik: rekening ? rekening.namaPemilik : '',
            nomorRekening: rekening ? rekening.nomorRekening : '',
            saldo: rekening ? rekening.saldo : 0
        },
        onSubmit: values => {
            if (rekening) {
                api.put(`/rekening/${rekening.id}`, values)
                    .then(response => {
                        onSave(response.data)
                    })
            } else {
                api.post('/rekening', values)
                    .then(response => {
                        onSave(response.data)
                    })
            }
        }
    })

    return (
        <form onSubmit={formik.handleSubmit}>
            <div>
                <label>Nama Pemilik</label>
                <input 
                    type="text" 
                    name="namaPemilik" 
                    value={formik.values.namaPemilik} 
                    onChange={formik.handleChange} 
                />
            </div>
            <div>
                <label>Nomor Rekening</label>
                <input 
                    type="text" 
                    name="nomorRekening" 
                    value={formik.values.nomorRekening} 
                    onChange={formik.handleChange} 
                />
            </div>
            <div>
                <label>Saldo</label>
                <input 
                    type="number" 
                    name="saldo" 
                    value={formik.values.saldo} 
                    onChange={formik.handleChange} 
                />
            </div>
            <button type="submit">Save</button>
        </form>
    )
}

export default RekeningForm
