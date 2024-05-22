import { useMutation } from "@tanstack/react-query";
import useGraphql from "../graphql";
import { fetchGraphQLWithToken, fetchUploadFile } from "../../../services/graphql-services";
import { toast } from "react-toastify";

function useCreateMovie(setFormData ,setSelectedImage , setSelectedFile , setMessage ){
    const {createMovie} = useGraphql()
    const {mutate} = useMutation({
        mutationFn:(input) => 
            fetchUploadFile(input),
            onSuccess:(data) => {
                toast.success('Tạo phim thành công')
            },
            onError:(error) =>{
                toast.error('Tạo phim thất bại')
            }
    })

    function onCreateMovie(movie , selectedFile) {
        try {
            const operation = {
                query:createMovie.query,
                variables: {
                  input:{
                    file: null,
                  ...movie
                  }
                },
              }
          
              const bodyFormData = new FormData()
              console.log(selectedFile)
              bodyFormData.append('operations', JSON.stringify(operation))
              bodyFormData.append('map', JSON.stringify({ 0: ['variables.input.file'] }))
              bodyFormData.append('0', selectedFile)
            
            mutate(bodyFormData);
            setFormData({
                title: '',
                director: '',
                cast: '',
                duration: '',
                genre: '',
                language: '',
                openingDay: '',
                rated: '',
                status: 'UPCOMING',
                story: '',
                trailer: '',
            });
            setSelectedImage(null);
            setSelectedFile(null);
            setMessage({ isShow: true, text: 'Thêm phim thành công', success: true });
        } catch (error) {
            console.log(error)
            console.log('Error during create Moive');
        }
    }

    return {onCreateMovie}
}

function useDeleteMovie(setNotification) {
    const {deleteMovie} = useGraphql()
    const {mutate} = useMutation({
        mutationFn:(input) => 
            fetchGraphQLWithToken(deleteMovie.query , {
                input:input
            }),
            onSuccess:(data) => {
                setNotification({ title: 'Thông báo', body: 'Đã xóa phim thành công!', footer: 'OK', status: 'success' });
                $('#deleteMovie').modal('hide');
                $('#notification').modal('show');
                toast.success('Xóa phim thành công')
            },
            onError:(error) =>{
                setNotification({ title: 'Thông báo', body: 'Phim đã có lịch chiếu, chỉ có thể xóa phim chưa có lịch chiếu!', footer: 'OK', status: 'danger' });
                $('#deleteMovie').modal('hide');
                $('#notification').modal('show');
                toast.error('Xóa phim thất bại')
            }
    })

    function onDeleteMovie(id){
        try {
            mutate(id);
        } catch (error) {
            console.log('Error during delete');
        }
    }

    return {onDeleteMovie}
}

function useCreateShowTime(generateTicket , setShowtimeData, setTimeStartAddShowtime,setIsConflict,setMessageAddShowtime) {
    const {createShowTime} = useGraphql()
    const {mutate} = useMutation({
        mutationFn: (input) => 
            fetchGraphQLWithToken(createShowTime.query,{
                input:input
            }),
            onSuccess: (data) =>{
                toast.success('Tạo suất chiếu thành công')
                console.log(data)
                generateTicket(data?.CreateShowTime?.id);
                setShowtimeData({});
                setTimeStartAddShowtime('');
                setIsConflict(true);
                setMessageAddShowtime({ isShow: false, text: '', success: false, item: null });
            },
            onError: (data) =>{
                toast.error('Tạo suất chiếu không thành công thời gian không phù hợp')
            }
    })

    function onCreateShowTime({startAt , endAt , movieId , roomId}){
        try{
            const start = startAt + 'Z'
            const end = endAt + 'Z'
            mutate({startAt:start , endAt:end , movieId , roomId})
        }catch (error) {
            console.log('Error during create');
        }
    }

    return {onCreateShowTime}
}

function useGenerateTickets(){
    const {generateTickets} = useGraphql()
    const {mutate} = useMutation({
        mutationFn: (input) => 
            fetchGraphQLWithToken(generateTickets.query,{
                input:input
            }),
            onSuccess: (data) =>{
                
            },
            onError: (data) =>{

            }
    })

    function onGenerateTickets({showTimeID , price}){
        try{
            mutate({showTimeID , price})
        }catch (error) {
            console.log('Error during generate');
        }
    }

    return {onGenerateTickets}
}

export {useCreateMovie , useDeleteMovie , useCreateShowTime,useGenerateTickets}