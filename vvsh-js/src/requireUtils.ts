// Courtercy of https://github.com/floatdrop/require-from-string/blob/master/index.js

'use strict';

import Module from "module"
var path = require('path');

export function requireFromString(codeBuffer: Buffer, id?: string, filename?: any, opts?: any) {

    const code = codeBuffer.toString()

	if (typeof filename === 'object') {
		opts = filename;
		filename = undefined;
	}

	opts = opts || {};
	filename = filename || '';

	opts.appendPaths = opts.appendPaths || [];
	opts.prependPaths = opts.prependPaths || [];

	if (typeof code !== 'string') {
		throw new Error('code must be a string, not ' + typeof code);
	}

	var paths = (Module as any)._nodeModulePaths(path.dirname(filename));

	var parent = module.parent;
    var m = new Module(filename, parent || undefined);
    
    id && (m.id = id)
	m.filename = filename;
	m.paths = [].concat(opts.prependPaths).concat(paths).concat(opts.appendPaths);
    (m as any)._compile(code, filename);

	var exports = m.exports;
	parent && parent.children && parent.children.splice(parent.children.indexOf(m), 1);

	return exports;
};

export function createModule(code: string, id: string, filename?: any, opts?: any){
	if (typeof filename === 'object') {
		opts = filename;
		filename = undefined;
	}

	opts = opts || {};
	filename = filename || '';

	opts.appendPaths = opts.appendPaths || [];
	opts.prependPaths = opts.prependPaths || [];

	if (typeof code !== 'string') {
		throw new Error('code must be a string, not ' + typeof code);
	}

	var paths = (Module as any)._nodeModulePaths(path.dirname(filename));

	var parent = module.parent;
    var m = new Module(filename, parent || undefined);
    
    id && (m.id = id)
	m.filename = m.id;
	m.paths = [].concat(opts.prependPaths).concat(paths).concat(opts.appendPaths);
	(m as any)._compile(code, filename);
	m.loaded = true
	return m
}


const ORIGIN_REQUIRE = require

export function useElRequier(moduleList: { [index: string]: Module}): NodeRequire{

	const f = (id: string) => {
		
		if (moduleList[id]){
			return moduleList[id]
		}
		return ORIGIN_REQUIRE(id)
	}
	return Object.assign(f, {
		cache: require.cache,
		extensions: require.extensions,
		resolve: require.resolve,
		main: require.main
	})

}

const INSIDE: { [index: string]: string } = {
	// "cmd-type": "https://edwinjhlee.github.io/cmd-type/index.html"
}

export function getNameUrl(e: string){
	const sep = e.indexOf(" ")
	if (sep < 0) {
		const name = e
		const url = INSIDE[name]
		return url ? { name: e, url: INSIDE[name] }: null
	}
	return {
		name: e.slice(0, sep).trim(),
		url: e.slice(sep+1).trim()
	}
}

import * as httpcache from "./http-cache"

//+use cmd-type https://edwinjhlee.github.io/cmd-type/cmd-type.js
export async function inject(code: string){

	const f = "//+use"
	const lines = code.split("\n")
		.filter(e => e.startsWith(f))
		.map(e => e.slice(f.length).trim())
		.filter(e => e.length !== 0)
	
	const ret = {} as { [index: string]: Module }
	for (let e of lines) {
		
		const w = getNameUrl(e)
		if (null === w) {
			throw new Error(`Lib not found: e`)
		}

        const res = await httpcache.get(w.url)
        if (res === false) {
            throw new Error("httpcache.fetch error")
        }
		ret[w.name] = requireFromString(res.code, w.name)
	}

	return await useElRequier(ret)
}

