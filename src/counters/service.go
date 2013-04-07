package counters

func NewCountRequest(name string) CountRequest {
  return CountRequest{&name, nil}
}

func NewCountResponse(count int32) CountResponse {
  return CountResponse{&count, nil}
}

// Sets the count of the current response.
func (this *CountResponse) SetCount(count int32) {
  // This feels so dirty...
  this.Count = &count
}

type CountClient struct {
  // The underlying service object.
  service CountersService
}

func NewCountersService(addr string) (CountClient, error) {
  service, err := DialCountersService(addr)

  if err != nil {
    return CountClient{nil}, err
  }

  return CountClient{service}, nil
}

func (this *CountClient) Increment(name string) int32 {
  request := NewCountRequest(name)
  response := CountResponse{}
  this.service.Increment(&request, &response)

  return response.GetCount()
}

func (this *CountClient) Decrement(name string) int32 {
  request := NewCountRequest(name)
  response := CountResponse{}
  this.service.Decrement(&request, &response)

  return response.GetCount()
}

func (this *CountClient) Get(name string) int32 {
  request := NewCountRequest(name)
  response := CountResponse{}
  this.service.Get(&request, &response)

  return response.GetCount()
}
