import RekeningForm from '../../components/RekeningForm'

const NewRekening = () => {
    return (
        <div>
            <h1>Tambah Rekening Baru</h1>
            <RekeningForm onSave={() => window.location.href = '/'} />
        </div>
    )
}

export default NewRekening
