import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { URL } from './init'
import { Xlsx } from './xlsx';

@Injectable({
  providedIn: 'root'
})
export class XlsxService {
  xlsxes: Xlsx[] = []
  constructor(
    private http: HttpClient,
  ) {
    this.load()
  }
  load() {
    this.http.get<Xlsx[]>(`${URL}/api/xlsxes`)
      .subscribe((xs: Xlsx[]) => {
        this.xlsxes = xs
      })
  }
  post(x: Xlsx) {
    if (x) {
      this.http.post<Xlsx>(`${URL}/api/xlsxes`, x)
        .subscribe((nx: Xlsx) => this.xlsxes.push(nx))
    }
  }
  async put(x: Xlsx) {
    const nx = await this.http.put<Xlsx>(`${URL}/api/xlsxes/${x.id}`, x).toPromise()
    Object.assign(x, nx)
    return x
  }
}
