import { NgZone } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { find, pull } from 'lodash'

import { Obj } from './obj';
import { URL } from './init'
import { Paging } from './paging';

export abstract class ObjService<T extends Obj> {
  private datas: T[] = []
  private txt: string = ''
  private tags: string[] = []
  private page = 0
  constructor(
    protected http: HttpClient,
    protected ngZone: NgZone,
  ) { }

  abstract path(): string

  get url() {
    return `${URL}${this.path()}`
  }

  get all$() {
    return this.http.get<T[]>(`${this.url}`, { params: this.params })
  }

  obj$(id: string) {
    return this.http.get<T>(`${this.url}/${id}`)
  }

  get paging$() {
    return this.http.get<Paging<T>>(`${this.url}`, { params: this.params })
  }

  get nextPaging$() {
    this.page += 1
    return this.paging$
  }

  has(datas: T[]): boolean {
    if (datas) {
      this.datas = datas
      return datas.length > 0
    }
    return false;
  }

  async save(o: T) {
    if (o.id) {
      const t = await this.http.put<T>(`${this.url}/${o.id}`, o).toPromise()
      const old = find(this.datas, (d) => d.id == t.id)
      if (old) {
        Object.assign(old, t)
      }
      return Object.assign(o, t)
    } else {
      const t = await this.http.post<T>(`${this.url}`, o).toPromise()
      this.datas.push(t)
      return t
    }
  }

  async delete(o: T) {
    await this.http.delete(`${this.url}/${o.id}`).toPromise()
    pull(this.datas, o)
    this.ngZone.run(() => false)
  }

  reset(t: T) {
    const o = find(this.datas, i => i.id == t.id)
    if (o) {
      Object.assign(o, t)
    }
  }

  search(txt: string) {
    this.txt = txt
    return this.all$
  }

  select(tags: string[]) {
    this.tags = tags
    return this.all$
  }

  searchPaging(txt: string) {
    this.page = 0
    this.txt = txt
    return this.paging$
  }

  selectPaging(tags: string[]) {
    this.page = 0
    this.tags = tags
    return this.paging$
  }


  private get params() {
    const params = {}
    if (this.txt) {
      params['search'] = this.txt
    }
    if (this.tags && this.tags.length > 0) {
      params['tags'] = JSON.stringify(this.tags)
    }
    if (this.page > 0) {
      params['page'] = this.page
    }
    return params
  }
}
