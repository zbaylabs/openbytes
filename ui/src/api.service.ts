//import { AttendantsPromiseClient } from './sdk/attendant_grpc_web_pb';
import { CapturesClient } from '../../api/js/capture_grpc_web_pb';

export class ApiService {
    // opts = { 'streamInterceptors': [new StreamInterceptor()] };;
    apiUrl = 'http://' + window.location.host
    // apiUrl = 'https://www.iyou.city:8443'

    // attendantClient = new AttendantsPromiseClient(this.apiUrl)
    // commodityClient = new CommoditiesPromiseClient(environment.apiUrl, null, this.opts);
    // userClient = new UsersPromiseClient(environment.apiUrl);
    // couponClient = new CouponsPromiseClient(environment.apiUrl);
    captureClient = new CapturesClient(this.apiUrl);
    // addressClient = new AddressesPromiseClient(environment.apiUrl);
    // messageClient = new MessagesPromiseClient(environment.apiUrl);
    // accountClient = new AccountsPromiseClient(environment.apiUrl);
    // commentClient = new CommentsPromiseClient(environment.apiUrl);
    // memoClient = new MemosPromiseClient(environment.apiUrl);

    //constructor() { }
    hospitals: any[] = [];
}

export const apiService = new ApiService();