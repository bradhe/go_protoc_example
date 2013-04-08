package counters

// A client wrapper for interacting with the CountersService service over RPC.
type CountClient struct {
  // The underlying service object.
  service CountersService
}

// Creates a new CountRequest object with the supplied name.
func NewCountRequest(name string) CountRequest {
  return CountRequest{&name, nil}
}

// Creates a new CountResponse object with the supplied count.
func NewCountResponse(count int32) CountResponse {
  return CountResponse{&count, nil}
}

// Sets the Count value attribute on the CountResponse object.
func (this *CountResponse) SetCount(count int32) {
  // This feels so dirty...
  this.Count = &count
}

// Creates a new client for the CountersService service.
func Dial(addr string) (CountClient, error) {
  service, err := DialCountersService(addr)

  if err != nil {
    return CountClient{nil}, err
  }

  return CountClient{service}, nil
}

// Increments the counter with the supplied name.
func (this *CountClient) Increment(name string) int32 {
  request := NewCountRequest(name)
  response := CountResponse{}
  this.service.Increment(&request, &response)

  return response.GetCount()
}

// Decrements the counter with the supplied name.
func (this *CountClient) Decrement(name string) int32 {
  request := NewCountRequest(name)
  response := CountResponse{}
  this.service.Decrement(&request, &response)

  return response.GetCount()
}

// Gets the counter with the supplied name.
func (this *CountClient) Get(name string) int32 {
  request := NewCountRequest(name)
  response := CountResponse{}
  this.service.Get(&request, &response)

  return response.GetCount()
}
