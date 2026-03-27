from rest_framework import status
from rest_framework.decorators import api_view
from rest_framework.response import Response

from .serializers import TaskSerializer

# In-memory store for simplicity
_tasks = {}
_next_id = 1


@api_view(["GET"])
def health(request):
    return Response({"status": "ok"})


@api_view(["GET", "POST"])
def task_list(request):
    global _next_id

    if request.method == "GET":
        return Response(list(_tasks.values()))

    serializer = TaskSerializer(data=request.data)
    serializer.is_valid(raise_exception=True)
    task = {
        "id": _next_id,
        "title": serializer.validated_data["title"],
        "description": serializer.validated_data.get("description", ""),
        "done": False,
    }
    _tasks[_next_id] = task
    _next_id += 1
    return Response(task, status=status.HTTP_201_CREATED)


@api_view(["GET", "DELETE"])
def task_detail(request, task_id):
    task = _tasks.get(task_id)
    if task is None:
        return Response(status=status.HTTP_404_NOT_FOUND)

    if request.method == "DELETE":
        del _tasks[task_id]
        return Response(status=status.HTTP_204_NO_CONTENT)

    return Response(task)
