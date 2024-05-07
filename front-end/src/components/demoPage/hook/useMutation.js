import { useGraphql } from "../graphql/index";
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { fetchGraphQL } from '../../../services/graphql-services'

function useDemoMutation() {
  const queryClient = useQueryClient()
  const { createJobTitle , queryKey } = useGraphql();
  const { mutate , isSuccess } = useMutation({
    mutationKey: [queryKey],
    mutationFn: (newTodo) =>
      fetchGraphQL(createJobTitle.query, {
        input: newTodo,
      }),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: [queryKey] }),
  })

  function onSubmit(value) {
     mutate(value)
  }
  return {
    isSuccess,
    onSubmit,
  };
}

export default useDemoMutation