/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntClassdate,
    EntClassdateFromJSON,
    EntClassdateFromJSONTyped,
    EntClassdateToJSON,
    EntClassroom,
    EntClassroomFromJSON,
    EntClassroomFromJSONTyped,
    EntClassroomToJSON,
    EntClasstime,
    EntClasstimeFromJSON,
    EntClasstimeFromJSONTyped,
    EntClasstimeToJSON,
    EntInstructorInfo,
    EntInstructorInfoFromJSON,
    EntInstructorInfoFromJSONTyped,
    EntInstructorInfoToJSON,
    EntSubject,
    EntSubjectFromJSON,
    EntSubjectFromJSONTyped,
    EntSubjectToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCourseclassEdges
 */
export interface EntCourseclassEdges {
    /**
     * 
     * @type {EntClassdate}
     * @memberof EntCourseclassEdges
     */
    classdate?: EntClassdate;
    /**
     * 
     * @type {EntClassroom}
     * @memberof EntCourseclassEdges
     */
    classroom?: EntClassroom;
    /**
     * 
     * @type {EntClasstime}
     * @memberof EntCourseclassEdges
     */
    classtime?: EntClasstime;
    /**
     * 
     * @type {EntInstructorInfo}
     * @memberof EntCourseclassEdges
     */
    instructorInfo?: EntInstructorInfo;
    /**
     * 
     * @type {EntSubject}
     * @memberof EntCourseclassEdges
     */
    subject?: EntSubject;
}

export function EntCourseclassEdgesFromJSON(json: any): EntCourseclassEdges {
    return EntCourseclassEdgesFromJSONTyped(json, false);
}

export function EntCourseclassEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCourseclassEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'classdate': !exists(json, 'Classdate') ? undefined : EntClassdateFromJSON(json['Classdate']),
        'classroom': !exists(json, 'Classroom') ? undefined : EntClassroomFromJSON(json['Classroom']),
        'classtime': !exists(json, 'Classtime') ? undefined : EntClasstimeFromJSON(json['Classtime']),
        'instructorInfo': !exists(json, 'InstructorInfo') ? undefined : EntInstructorInfoFromJSON(json['InstructorInfo']),
        'subject': !exists(json, 'Subject') ? undefined : EntSubjectFromJSON(json['Subject']),
    };
}

export function EntCourseclassEdgesToJSON(value?: EntCourseclassEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'classdate': EntClassdateToJSON(value.classdate),
        'classroom': EntClassroomToJSON(value.classroom),
        'classtime': EntClasstimeToJSON(value.classtime),
        'instructorInfo': EntInstructorInfoToJSON(value.instructorInfo),
        'subject': EntSubjectToJSON(value.subject),
    };
}


