docker_build('maxhorstmann/todo', '.', dockerfile = 'Dockerfile')

k8s_yaml('services.yaml')
k8s_resource('sql', port_forwards="1433")
k8s_resource('todo', port_forwards="1234")
k8s_resource('grpcui', port_forwards="8080")

